FROM golang:1.19-alpine as dev

WORKDIR /home/rest-api
COPY Makefile .
RUN apk update && apk add --no-cache build-base && apk add curl
RUN make tools

# Image for build
FROM golang:1.19-alpine as build
WORKDIR /home/rest-api
RUN apk update && apk add --no-cache build-base

COPY go.mod go.sum ./
RUN go mod download && go mod verify 
COPY . .

RUN make build
RUN chmod +x ./bin/matryer

# Image for production
FROM alpine:latest as prod
WORKDIR /home/rest-api

COPY --from=build /home/rest-api/bin/matryer /home/rest-api/
CMD ["./matryer"]