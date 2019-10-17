FROM openfaas/classic-watchdog:0.13.4 as watchdog
FROM alpine:3.9

RUN apk add python3-dev py3-lxml curl
RUN pip3 install bs4 requests

COPY dcurooms.py /dcurooms.py

RUN apk --no-cache add curl \
    && echo "Pulling watchdog binary from Github." \
    && curl -sSL https://github.com/openfaas/faas/releases/download/0.9.6/fwatchdog > /usr/bin/fwatchdog \
    && apk del curl --no-cache

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog


# Populate example here - i.e. "cat", "sha512sum" or "node index.js"
ENV fprocess="xargs python3 dcurooms.py"
# Set to true to see request in function logs
ENV write_debug="false"

EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]
