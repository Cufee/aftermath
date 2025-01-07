# Build javascript
FROM oven/bun:latest as builder-js

WORKDIR /workspace

COPY ./cmd/frontend/widget/package.json ./
RUN bun install

COPY ./cmd/frontend/widget/ ./

RUN bun run build

# Build app
FROM golang:1.23-bookworm as builder-go

ARG BRAND_FLAVOR=red
ENV BRAND_FLAVOR $BRAND_FLAVOR

WORKDIR /workspace

COPY go.mod go.sum ./
RUN --mount=type=cache,target=$GOPATH/pkg/mod go mod download

# Install templ
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY ./ ./

# build a fully standalone binary with zero dependencies
RUN --mount=type=cache,target=$GOPATH/pkg/mod go generate ./internal/assets
RUN --mount=type=cache,target=$GOPATH/pkg/mod go generate ./cmd/frontend/assets/generate

RUN templ generate

# Copy js assets
COPY --from=builder-js /workspace/dist/components/* ./cmd/frontend/public/js/widget/

RUN --mount=type=cache,target=$GOPATH/pkg/mod CGO_ENABLED=1 GOOS=linux go build -ldflags='-s -w' -trimpath -o /bin/aftermath .

# Make a scratch container with required files and binary
FROM debian:stable-slim

ENV TZ=Etc/UTC
ENV ZONEINFO=/zoneinfo.zip
COPY --from=builder-go /bin/aftermath /usr/bin/aftermath
COPY --from=builder-go /usr/local/go/lib/time/zoneinfo.zip /
COPY --from=builder-go /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "aftermath" ]
