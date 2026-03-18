ARG GO_VERSION=1.24.2

FROM golang:${GO_VERSION}-alpine AS build
LABEL "language"="go"
LABEL "framework"="echo"

WORKDIR /app

RUN apk update -qq && apk add --no-cache git

COPY . .

RUN go build -o api main.go wire_gen.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=build /app/api /app/api

EXPOSE 8002

ENTRYPOINT ["/app/api"]