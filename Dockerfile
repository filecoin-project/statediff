FROM node:14-alpine AS js
WORKDIR /usr/src/app
COPY ./npm/app .
RUN npm install

FROM golang:alpine AS builder
RUN apk update
RUN apk upgrade
RUN apk add --update gcc>=9.3.0 g++>=9.3.0 alpine-sdk

WORKDIR /go/src/app/

COPY . .
COPY --from=js /usr/src/app ./npm/app
# Fetch dependencies.
RUN go get -d -v ./... && go get github.com/go-bindata/go-bindata/go-bindata
RUN BUILD_DIR=$(pwd); GOZSTD_VER=$(cat go.mod | fgrep github.com/valyala/gozstd | awk '{print $2}'); go get -d github.com/valyala/gozstd@${GOZSTD_VER}; cd ${GOPATH}/pkg/mod/github.com/valyala/gozstd@${GOZSTD_VER}; if [[ ! -f _rebuilt ]]; then chmod -R +w .; make -j8 clean; make -j8 libzstd.a; touch _rebuilt; fi; cd ${BUILD_DIR}
RUN go generate ./...
RUN go build -o stateexplorer ./cmd/stateexplorer

FROM alpine
# Copy our static executable.
COPY --from=builder /go/src/app/stateexplorer /stateexplorer
ENV REALIZED_CACHE_SIZE 32768
ENV BLOCK_CACHE_SIZE 32768
ENTRYPOINT ["/stateexplorer"]
