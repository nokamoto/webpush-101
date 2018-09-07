#!/bin/sh

set -ex

cd `dirname $0`

VERSION=`cat VERSION`

time docker build -t nokamoto/webpush-101.go.webpush:${VERSION} -f docker/webpush/Dockerfile .
