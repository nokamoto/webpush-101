#!/bin/bash

set -ex

cd `dirname $0`

VERSION=`cat VERSION | tr -d '\n'`

time docker build -t nokamoto13/webpush-101.go.webpush:${VERSION} -f docker/webpush/Dockerfile .
time docker build -t nokamoto13/webpush-101.go.push-subscription:${VERSION} -f docker/push-subscription/Dockerfile .

if [ "$1" = "--no-sbt-package" ]
then
    echo no sbt universal:packageZipTarball
else
    (cd webpush-scala/front && sbt universal:packageZipTarball)
    (cd webpush-scala/webpush && sbt universal:packageZipTarball)
fi

time docker build -t nokamoto13/webpush-101.scala.front:${VERSION} -f docker/front/Dockerfile --build-arg VERSION=${VERSION} .
time docker build -t nokamoto13/webpush-101.scala.webpush:${VERSION} -f docker/scala.webpush/Dockerfile --build-arg VERSION=${VERSION} .

docker push nokamoto13/webpush-101.go.webpush:${VERSION}
docker push nokamoto13/webpush-101.go.push-subscription:${VERSION}
docker push nokamoto13/webpush-101.scala.front:${VERSION}
docker push nokamoto13/webpush-101.scala.webpush:${VERSION}
