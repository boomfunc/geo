# Prepare app code (some compiling and other)
FROM golang:alpine as app

ADD . /boomfunc/app

RUN set -ex \
	&& apk add --update --no-cache git \
	&& cd /boomfunc/app \
	&& go get -v -d \
	&& go build geoip.go

# Final container, copy from builders
# Get pre-compiled base
FROM boomfunc/base:latest as base

COPY --from=app /boomfunc/app /boomfunc/app
