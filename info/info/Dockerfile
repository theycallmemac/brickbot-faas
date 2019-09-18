FROM openfaas/classic-watchdog:0.15.4 as watchdog

FROM alpine:3.9

RUN apk add bash curl jq 

RUN curl -s https://gist.githubusercontent.com/theycallmemac/5b401a0475011206a4944f531719c643/raw/b0f30122804de2908e71608db21528167e0fa43f/botinfo > botinfo
RUN curl -s https://raw.githubusercontent.com/theycallmemac/brickbot/master/package.json > package.json

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

# Populate example here - i.e. "cat", "sha512sum" or "node index.js"
ENV fprocess="bash botinfo"

# Set to true to see request in function logs
ENV write_debug="false"

EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]
