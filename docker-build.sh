#!/bin/bash

set -ex

cd `dirname $0`

VERSION=`cat VERSION | tr -d '\n'`

time docker build -t nokamoto/webpush-101.go.webpush:${VERSION} -f docker/webpush/Dockerfile .

if [ "$1" = "--no-sbt-package" ]
then
    echo no sbt universal:packageZipTarball
else
    (cd webpush-scala/front && sbt universal:packageZipTarball)
fi

time docker build -t nokamoto/webpush-101.scala.front:${VERSION} -f docker/front/Dockerfile --build-arg VERSION=${VERSION} .
