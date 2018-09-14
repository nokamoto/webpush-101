FROM golang:1.11.0-alpine3.8 AS build

RUN apk add git

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/nokamoto/webpush-101

COPY Gopkg.lock .
COPY Gopkg.toml .
COPY webpush-go webpush-go

RUN dep ensure -vendor-only=true

RUN go install ./webpush-go/push-subscription

FROM alpine:3.8

RUN apk update
RUN apk add --no-cache ca-certificates

COPY --from=build /go/bin/push-subscription /usr/local/bin/puhs-subscription

ENV GRPC_PORT=8000

ENTRYPOINT [ "push-subscription" ]

CMD [ "-h" ]
