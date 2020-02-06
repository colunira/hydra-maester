FROM golang:1.12.9 as builder

WORKDIR /go/src/github.com/ory/hydra-maester/synchronization
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./sync.go

RUN mkdir /app && mv ./main /app/main

FROM alpine
WORKDIR /app

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add --no-cache bash

COPY --from=builder /app/main /app
COPY ./entrypoint.sh .

ENTRYPOINT ["/bin/bash", "-c", "./entrypoint.sh"]