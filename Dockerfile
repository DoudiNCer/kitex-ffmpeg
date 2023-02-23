FROM alpine:latest
MAINTAINER doudincer<doudi@outlook.lv>
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.bfsu.edu.cn/g' /etc/apk/repositories &&  \
    apk update && apk upgrade &&  \
    apk --no-cache add yasm ffmpeg
ENV VERSION 1.1
WORKDIR /apps
ENV KITEX_LOG_DIR="/apps/log"
COPY output/bin/KitexFfmpeg ./KitexFfmpeg
EXPOSE 9427
CMD ["/apps/KitexFfmpeg"]
