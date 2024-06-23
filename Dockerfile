FROM golang:1.22.3-bookworm as builder

WORKDIR /workspace

COPY go.mod go.sum ./
RUN --mount=type=cache,target=$GOPATH/pkg/mod go mod download

COPY ./ ./

# build a fully standalone binary with zero dependencies
RUN --mount=type=cache,target=$GOPATH/pkg/mod GOOS=linux go build -a -installsuffix cgo -o /bin/aftermath .

# Make a scratch container with required files and binary
FROM debian:stable-slim

ENV TZ=Europe/Berlin
COPY --from=builder /bin/aftermath /usr/bin/aftermath

ENTRYPOINT [ "aftermath" ]
