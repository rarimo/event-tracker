FROM golang:1.19-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/rarimo/identity/event-tracker
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/event-tracker /go/src/gitlab.com/rarimo/identity/event-tracker


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/event-tracker /usr/local/bin/event-tracker
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["event-tracker"]
