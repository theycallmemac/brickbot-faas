FROM openfaas/classic-watchdog:0.13.4 as watchdog

FROM alpine:3.9

RUN apk add bash curl

RUN curl -s https://gist.githubusercontent.com/theycallmemac/f66b0afeca215df97869dd28612bea74/raw/83d570e6a95122068932901ef55c58b881726e1d/flip.sh > /tmp/flip.sh

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

# Populate example here - i.e. "cat", "sha512sum" or "node index.js"
ENV fprocess="bash /tmp/flip.sh"
# Set to true to see request in function logs
ENV write_debug="false"

EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]
