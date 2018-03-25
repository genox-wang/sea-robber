FROM busybox

RUN echo "Asia/Shanghai" > /etc/timezone

ADD /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/

WORKDIR /app

EXPOSE 8000

CMD [ "./sea-robber-api" ]