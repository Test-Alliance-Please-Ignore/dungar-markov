FROM golang:1.21-bookworm AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=mod -o /out/dungar -ldflags="-s -w" ./cmd/dungar

FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    useradd --system --create-home --home-dir /app dungar

WORKDIR /app

COPY --from=build /out/dungar /usr/local/bin/dungar

USER dungar

ENTRYPOINT ["/usr/local/bin/dungar"]
CMD ["run"]
