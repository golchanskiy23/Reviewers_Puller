FROM golang:1.24

WORKDIR /app

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        tzdata \
        make \
        postgresql-client \
        curl && \
    rm -rf /var/lib/apt/lists/*

RUN apt-get update && apt-get install -y tini && rm -rf /var/lib/apt/lists/*

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
RUN go get github.com/prometheus/client_golang@latest

COPY . .

RUN go mod download

EXPOSE 8080

ENTRYPOINT ["/usr/bin/tini", "--"]
CMD ["make", "start"]
