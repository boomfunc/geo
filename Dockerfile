FROM boomfunc/app:dev

RUN mkdir -p /bmp

WORKDIR /bmp

ADD . /bmp
