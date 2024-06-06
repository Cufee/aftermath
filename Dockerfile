FROM golang:1.22.3-alpine as builder

WORKDIR /workspace

# add go modules lockfiles
COPY go.mod go.sum ./
RUN go mod download

# prefetch the binaries, so that they will be cached and not downloaded on each change
RUN go run github.com/steebchen/prisma-client-go prefetch

COPY ./ ./
# generate the Prisma Client Go client
RUN go generate ./...

# build a fully standalone binary with zero dependencies
RUN CGO_ENABLED=0 GOOS=linux go build -o app .


# <!> railway.app does not really have a way to run a migration container on the volume, so we have to improvise here
# in order to run migrations, we need to have go and prisma installed
FROM golang:1.22.3-alpine

# set timezone
ENV TZ=Europe/Berlin

# copy final binary
COPY --from=builder /workspace/app /app

# copy migrations and schema
COPY --from=builder /workspace/internal/database/prisma/migrations /prisma/migrations
COPY --from=builder /workspace/internal/database/prisma/schema.prisma /prisma/schema.prisma
COPY --from=builder /workspace/docker-entrypoint.sh /docker-entrypoint.sh
COPY --from=builder /workspace/go.mod /go.mod
COPY --from=builder /workspace/go.sum /go.sum

# install prisma and prefetch binaries
RUN go install github.com/steebchen/prisma-client-go
RUN go run github.com/steebchen/prisma-client-go prefetch
RUN chmod +x /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]
