FROM node:14-alpine AS js
WORKDIR /usr/src/app
COPY ./npm/app .
RUN npm install

FROM golang:alpine AS builder
RUN apk update
RUN apk upgrade
RUN apk add --update gcc>=9.3.0 g++>=9.3.0

WORKDIR /go/src/app/

COPY . .
COPY --from=js /usr/src/app ./npm/app
# Fetch dependencies.
RUN go get -d -v ./...
RUN go generate ./...
RUN go build -o stateexplorer ./cmd/stateexplorer

FROM alpine
# Copy our static executable.
COPY --from=builder /go/src/app/stateexplorer /stateexplorer
ENV REALIZED_CACHE_SIZE 32768
ENTRYPOINT ["/stateexplorer"]
