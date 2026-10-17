FROM alpine:3.20
RUN apk add --no-cache bash curl
COPY . /cron/
RUN chmod +x /cron/*.sh /cron/jobs/*.sh && crontab /cron/crontab
CMD ["sh", "-c", "export -p > /cron/env.sh && exec crond -f -l 8"]
