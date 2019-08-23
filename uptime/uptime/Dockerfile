FROM openfaas/classic-watchdog:0.15.4 as watchdog

FROM alpine:3.9

RUN apk add nodejs curl

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

RUN curl -s https://gist.githubusercontent.com/theycallmemac/36a0bfa7455c74dd50a76d2190478e91/raw/f8d9d5d78e294e5426ff695127f4f15d09424843/discordbot-uptime.js > uptime.js

# Populate example here - i.e. "cat", "sha512sum" or "node index.js"
ENV fprocess="xargs node uptime.js"
# Set to true to see request in function logs
ENV write_debug="false"

EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]
