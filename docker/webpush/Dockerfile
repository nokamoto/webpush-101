FROM golang:1.11.0-alpine3.8 AS build

RUN apk add git

RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/nokamoto/webpush-101

COPY Gopkg.lock .
COPY Gopkg.toml .
COPY webpush-go webpush-go

RUN dep ensure -vendor-only=true

RUN go install ./webpush-go/webpush

FROM alpine:3.8

RUN apk update
RUN apk add --no-cache ca-certificates

COPY --from=build /go/bin/webpush /usr/local/bin/webpush

ENV GRPC_PORT=8000

ENTRYPOINT [ "webpush" ]

CMD [ "-h" ]
