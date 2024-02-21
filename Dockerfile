# syntax=docker/dockerfile:1

# local env
FROM golang:1.21-alpine3.19 AS local

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.49.0

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENTRYPOINT ["air"]


# build
FROM golang:1.21-alpine3.19 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /simple-kanban .


# deployable
FROM alpine:3.19

WORKDIR /

COPY --from=build /simple-kanban /simple-kanban

EXPOSE 3000

ENTRYPOINT ["/server"]
