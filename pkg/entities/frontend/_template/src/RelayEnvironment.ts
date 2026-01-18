import { Environment, Network, RecordSource, Store, FetchFunction } from 'relay-runtime';

const HTTP_ENDPOINT = import.meta.env.VITE_GRAPHQL_ENDPOINT || '/graphql';

const fetchFn: FetchFunction = async (request, variables) => {
  const response = await fetch(HTTP_ENDPOINT, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      query: request.text,
      variables,
    }),
  });

  const json = await response.json();

  if (Array.isArray(json.errors)) {
    console.error(json.errors);
    throw new Error(
      `Error fetching GraphQL query '${request.name}' with variables '${JSON.stringify(variables)}': ${JSON.stringify(json.errors)}`,
    );
  }

  return json;
};

function createRelayEnvironment() {
  return new Environment({
    network: Network.create(fetchFn),
    store: new Store(new RecordSource()),
  });
}

export const RelayEnvironment = createRelayEnvironment();
