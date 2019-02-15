# Building stage
FROM golang:1.11-alpine as build

RUN mkdir -p /go/src/github.com/nylo-andry/monitor360

WORKDIR /go/src/github.com/nylo-andry/monitor360

COPY vendor vendor
COPY cmd cmd
COPY http http
COPY monitor360.go .

RUN GOOS=linux go build -o ./bin/app cmd/server/main.go

# Running stage
FROM alpine:latest

RUN apk update 
RUN apk add ca-certificates 
RUN rm -rf /var/cache/apk/*

WORKDIR /service
COPY --from=build /go/src/github.com/nylo-andry/monitor360/bin/app /service/monitor360

CMD ["./monitor360"]