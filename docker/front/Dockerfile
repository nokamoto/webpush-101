FROM openjdk:8u181-jre-slim

ARG VERSION
ENV VERSION=${VERSION}

COPY webpush-scala/front/target/universal/front-${VERSION}.tgz .

RUN tar -zxvf front-${VERSION}.tgz

RUN mv front-${VERSION} front

RUN rm front/conf/application.conf

ENTRYPOINT [ "front/bin/front" ]

CMD [ "-h" ]
