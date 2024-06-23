FROM golang:1.22.3-alpine as builder

WORKDIR /workspace

COPY go.mod go.sum ./
RUN --mount=type=cache,target=$GOPATH/pkg/mod go mod download

COPY ./ ./
RUN --mount=type=cache,target=$GOPATH/pkg/mod go generate ./...

# build a fully standalone binary with zero dependencies
RUN --mount=type=cache,target=$GOPATH/pkg/mod CGO_ENABLED=0 GOOS=linux go build -o app .

# Make a scratch container with required files and binary
FROM scratch

WORKDIR /app

ENV TZ=Europe/Berlin
ENV ZONEINFO=/zoneinfo.zip
COPY --from=builder /workspace/app /usr/bin/
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["app"]
