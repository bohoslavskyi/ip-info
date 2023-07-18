# ------------ Stage 0 - build env --------------
FROM golang:1.20-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download
COPY . ./

RUN go build -v -o server cmd/main.go
# ------------ end of Stage 0 ---------------

# ------------ Stage 1 - build --------------
FROM debian:buster-slim

ENV SERVER_PORT 80

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server
COPY .env .env

EXPOSE ${SERVER_PORT}
CMD ["/app/server"]
# ------------ end of Stage 1 --------------
