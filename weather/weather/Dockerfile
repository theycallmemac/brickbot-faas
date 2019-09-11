FROM openfaas/classic-watchdog:0.15.4 as watchdog

FROM alpine:3.9

RUN apk add curl bash

RUN curl -s https://gist.githubusercontent.com/theycallmemac/247a89ba667ae91df696dc6a4e29edd7/raw/34eb6671627790fe20cae28922797a02fccb83d5/weatherByLocation.sh > /tmp/weatherByLocation.sh

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

# Populate example here - i.e. "cat", "sha512sum" or "node index.js"
ENV fprocess="xargs bash /tmp/weatherByLocation.sh"
# Set to true to see request in function logs
ENV write_debug="false"

EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]
