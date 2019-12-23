FROM alpine:3.8
RUN apk add --no-cache bash curl jq pwgen
RUN echo "Pulling watchdog binary from Github." \
    && curl -sSL https://github.com/openfaas/faas/releases/download/0.9.6/fwatchdog > /usr/bin/fwatchdog \
    && chmod +x /usr/bin/fwatchdog \
    && cp /usr/bin/fwatchdog /
ADD select_function.sh /select_function.sh
ADD functions /functions
ENV fprocess="xargs bash /select_function.sh"
ENV write_debug="true"
EXPOSE 8080
HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1
CMD [ "fwatchdog" ]
