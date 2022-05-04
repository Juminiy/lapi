# Pysical Machine,Virtual Machine,Linux OS
FROM alpine:latest
MAINTAINER Chisato
ENV VERSION 1.0

WORKDIR /apps

COPY lapi /apps/lapi

COPY .env /apps/.env

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

ENV LANG C.UTF-8

EXPOSE 8000

ENTRYPOINT ["/apps/lapi"]
