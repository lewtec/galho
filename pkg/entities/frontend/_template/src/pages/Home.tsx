import { Suspense } from 'react';
import { useTranslation } from 'react-i18next';
import { graphql } from 'react-relay';

// This query will be compiled by relay-compiler
// Make sure your GraphQL schema has a 'viewer' query or adjust accordingly
// eslint-disable-next-line @typescript-eslint/no-unused-vars
const HomeQuery = graphql`
  query HomeQuery {
    viewer {
      id
    }
  }
`;

function HomeContent() {
  const { t } = useTranslation();

  // Uncomment this when you have a GraphQL server running
  // const data = useLazyLoadQuery(HomeQuery, {});

  return (
    <div className="space-y-8">
      <div className="hero bg-base-100 rounded-lg shadow-xl">
        <div className="hero-content py-16 text-center">
          <div className="max-w-md">
            <h1 className="text-5xl font-bold">{t('home.title')}</h1>
            <p className="py-6">{t('home.subtitle')}</p>
            <button className="btn btn-primary">Get Started</button>
          </div>
        </div>
      </div>

      <div className="card bg-base-100 shadow-xl">
        <div className="card-body">
          <h2 className="card-title">Lorem Ipsum</h2>
          <p>
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor
            incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud
            exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
          </p>
          <p>
            Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat
            nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui
            officia deserunt mollit anim id est laborum.
          </p>
        </div>
      </div>

      <div className="grid grid-cols-1 gap-4 md:grid-cols-3">
        <div className="card bg-base-100 shadow-xl">
          <div className="card-body">
            <h3 className="card-title">Feature 1</h3>
            <p>
              Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor
              incididunt ut labore.
            </p>
          </div>
        </div>
        <div className="card bg-base-100 shadow-xl">
          <div className="card-body">
            <h3 className="card-title">Feature 2</h3>
            <p>
              Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor
              incididunt ut labore.
            </p>
          </div>
        </div>
        <div className="card bg-base-100 shadow-xl">
          <div className="card-body">
            <h3 className="card-title">Feature 3</h3>
            <p>
              Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor
              incididunt ut labore.
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default function Home() {
  return (
    <Suspense fallback={<div className="loading loading-spinner loading-lg"></div>}>
      <HomeContent />
    </Suspense>
  );
}
