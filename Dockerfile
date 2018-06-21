FROM boomfunc/app:dev

RUN mkdir -p /bmp

WORKDIR /bmp

ADD . /bmp

# custom additional actions
ENV PYGEOIP_VERSION 0.3.2

RUN set -ex \
		&& curl https://bootstrap.pypa.io/get-pip.py -o /tmp/get-pip.py \
		&& python /tmp/get-pip.py \
		\
		&& pip install --no-cache-dir pygeoip==${PYGEOIP_VERSION} \
		\
		&& rm /tmp/get-pip.py
