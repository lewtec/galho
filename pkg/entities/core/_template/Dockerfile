FROM golang:1.25-alpine AS builder

WORKDIR /app

# Scaffolded apps may not have go.sum until the first tidy/download.
COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o app ./cmd/app

FROM scratch

WORKDIR /
COPY --from=builder /app/app .

ENTRYPOINT ["/app"]
