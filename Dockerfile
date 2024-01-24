FROM golang:1.21.5-alpine

WORKDIR /home/rest-api
COPY Makefile .
RUN apk update && apk add --no-cache build-base && apk add curl
RUN make tools