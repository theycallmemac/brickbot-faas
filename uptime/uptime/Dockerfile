FROM openfaas/classic-watchdog:0.15.4 as watchdog

FROM alpine:3.9

RUN apk add nodejs

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

COPY ./uptime.js /uptime.js

# Populate example here - i.e. "cat", "sha512sum" or "node index.js"
ENV fprocess="xargs node uptime.js"
# Set to true to see request in function logs
ENV write_debug="false"

EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]
