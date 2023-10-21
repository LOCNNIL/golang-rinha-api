## Build
FROM golang:alpine AS build

RUN apk update && apk add --no-cache git

WORKDIR /usr/local/go/src/github.com/LOCNNIL/golang-rinha-api

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY cmd/ ./cmd/
COPY app/ ./app/

RUN go build -mod=mod -o /rinha-api cmd/main.go

## Deploy
FROM scratch

WORKDIR /

COPY --from=build /rinha-api /rinha-api

EXPOSE 8081

ENTRYPOINT ["/rinha-api"]

