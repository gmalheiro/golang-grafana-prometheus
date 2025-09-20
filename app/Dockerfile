# Build stage
FROM golang:1.25-alpine3.22 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY cmd/api ./cmd/api
COPY internal ./internal
RUN go build -o server ./cmd/api

# Runtime stage
FROM alpine:3.22.1

WORKDIR /app
COPY --from=build /app/server .

ENTRYPOINT ["/app/server"]
