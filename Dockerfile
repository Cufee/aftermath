# Download localizations
FROM node:23 AS builder-node

ARG LOCALIZE_API_KEY
ENV LOCALIZE_API_KEY $LOCALIZE_API_KEY

WORKDIR /workspace

RUN npm install @tolgee/cli

COPY ./.tolgeerc ./

RUN npx tolgee pull --api-key "${LOCALIZE_API_KEY}" --states REVIEWED

# Build app
FROM golang:latest AS builder-go

ARG BRAND_FLAVOR=red
ENV BRAND_FLAVOR $BRAND_FLAVOR

WORKDIR /workspace

COPY go.mod go.sum ./
RUN --mount=type=cache,target=$GOPATH/pkg/mod go mod download

COPY ./ ./

# load localizations
COPY --from=builder-node /workspace/static/localization/ ./static/localization/

# generate static assets
RUN --mount=type=cache,target=$GOPATH/pkg/mod go generate ./internal/assets
RUN --mount=type=cache,target=$GOPATH/pkg/mod go generate ./cmd/frontend/assets/generate

# generate frontend
RUN go tool templ generate

# build a fully standalone binary with zero dependencies
RUN --mount=type=cache,target=$GOPATH/pkg/mod CGO_ENABLED=1 GOOS=linux go build -ldflags='-s -w' -trimpath -o /bin/aftermath .

# Make a scratch container with required files and binary
FROM scratch

ENV TZ=Etc/UTC
ENV ZONEINFO=/zoneinfo.zip
COPY --from=builder-go /bin/aftermath /usr/bin/aftermath
COPY --from=builder-go /usr/local/go/lib/time/zoneinfo.zip /
COPY --from=builder-go /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "aftermath" ]
