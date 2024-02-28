# Build and pack

FROM golang:1.22.0-alpine3.1 AS golang-builder

RUN go version 

ARG VERSION

COPY . /srv
WORKDIR /srv

RUN set -Eeux && go mod download && go mod verify

RUN GOOS=linux \ 
    GOARCH=amd64 \ 
    go build \ 
    -o app -ldflags=-X=main.version=${VERSION}  main.go

FROM alpine:3.19.1
WORKDIR /srv
COPY --from=golang-builder /srv/app .

EXPOSE 8080
ENTRYPOINT ["./app"]