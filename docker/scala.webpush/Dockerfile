FROM openjdk:8u181-jre-slim

ARG VERSION
ENV VERSION=${VERSION}

COPY webpush-scala/webpush/target/universal/webpush-${VERSION}.tgz .

RUN tar -zxvf webpush-${VERSION}.tgz

RUN mv webpush-${VERSION} webpush

ENTRYPOINT [ "webpush/bin/webpush" ]

CMD [ "-h" ]
