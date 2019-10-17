FROM openfaas/classic-watchdog:0.13.4 as watchdog

FROM alpine:3.9

RUN apk add bash

COPY coinflip.sh /bin/coinflip.sh 

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

# Populate example here - i.e. "cat", "sha512sum" or "node index.js"
ENV fprocess="bash /bin/coinflip.sh"
# Set to true to see request in function logs
ENV write_debug="false"

EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]
