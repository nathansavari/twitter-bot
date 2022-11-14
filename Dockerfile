FROM golang:1.18-alpine as build

WORKDIR /app

COPY . .

RUN env GOOS=linux GOARCH=amd64 go build

FROM alpine:latest

ARG BUILD_CONSUMER_KEY
ENV CONSUMER_KEY=BUILD_CONSUMER_KEY

ARG BUILD_ACCESS_TOKEN_KEY
ENV ACCESS_TOKEN_KEY=BUILD_ACCESS_TOKEN_KEY

ARG BUILD_CONSUMER_SECRET
ENV CONSUMER_SECRET=BUILD_CONSUMER_SECRET

ARG BUILD_ACCESS_TOKEN_SECRET
ENV ACCESS_TOKEN_SECRET=BUILD_ACCESS_TOKEN_SECRET

ENV ENV=prod

WORKDIR /app

RUN apk update
RUN apk upgrade
RUN apk add --no-cache bash

COPY  --from=build /app/twitter-bot /app/twitter-bot

RUN chmod +x /app/twitter-bot

ENTRYPOINT [ "/bin/bash", "-c", "/app/twitter-bot" ]
