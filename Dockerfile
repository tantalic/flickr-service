FROM gliderlabs/alpine
MAINTAINER Kevin Stock <kevinstock@tantalic.com>

RUN apk-install ca-certificates
ADD ./flickr-service /usr/local/bin/

ENTRYPOINT /usr/local/bin/flickr-service
EXPOSE 3000
