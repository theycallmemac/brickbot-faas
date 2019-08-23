FROM openfaas/classic-watchdog:0.13.4 as watchdog
FROM alpine:3.9

RUN apk add --no-cache  python3-dev py3-lxml curl
RUN pip3 install bs4 requests

RUN mkdir -p /home/app

RUN apk --no-cache add curl \
    && echo "Pulling watchdog binary from Github." \
    && curl -sSL https://github.com/openfaas/faas/releases/download/0.9.6/fwatchdog > /usr/bin/fwatchdog \
    && curl -s https://gist.githubusercontent.com/theycallmemac/43b8c9973b978fa7cb2769cc6f3d4cdc/raw/976fc353f50d7715e580e7a635aab18b53e4ab93/roomcheck.py > roomcheck.py \
    && apk del curl --no-cache

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog


# Populate example here - i.e. "cat", "sha512sum" or "node index.js"
ENV fprocess="xargs python3 roomcheck.py"
# Set to true to see request in function logs
ENV write_debug="false"

EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]
