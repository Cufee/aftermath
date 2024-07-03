FROM golang:1.22.3-bookworm as builder

WORKDIR /workspace

COPY go.mod go.sum ./
RUN --mount=type=cache,target=$GOPATH/pkg/mod go mod download

COPY ./ ./

# build a fully standalone binary with zero dependencies
RUN --mount=type=cache,target=$GOPATH/pkg/mod go generate main.go
RUN --mount=type=cache,target=$GOPATH/pkg/mod CGO_ENABLED=1 GOOS=linux go build -o /bin/aftermath .

# Make a scratch container with required files and binary
FROM debian:stable-slim

ENV TZ=Europe/Berlin
ENV ZONEINFO=/zoneinfo.zip
COPY --from=builder /bin/aftermath /usr/bin/aftermath
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "aftermath" ]
