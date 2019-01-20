FROM nginx:alpine
MAINTAINER Dan Richards <dan.m.richards@gmail.com>

RUN mkdir -p /var/log/nginx && mkdir -p /tmp/nginx && chown nginx /tmp/nginx

COPY docker/nginx/conf/nginx.conf /etc/nginx/conf.d/default.conf