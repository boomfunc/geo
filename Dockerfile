FROM boomfunc/base:latest

ADD . ${BMPROOT}/app

# custom additional actions
ENV PYGEOIP_VERSION 0.3.2

RUN set -ex \
		&& apk add --no-cache \
			python \
			py-pip \
		&& pip install --no-cache-dir pygeoip==${PYGEOIP_VERSION} \
		\
		&& rm -rf /var/cache/apk/* \
		&& apk del py-pip
