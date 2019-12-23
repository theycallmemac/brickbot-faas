FROM alpine:3.8
RUN apk add --no-cache bash python3 curl
RUN echo "Pulling watchdog binary from Github." \
    && curl -sSL https://github.com/openfaas/faas/releases/download/0.9.6/fwatchdog > /usr/bin/fwatchdog \
    && chmod +x /usr/bin/fwatchdog \
    && cp /usr/bin/fwatchdog /
ADD select_function.sh /select_function.sh
ADD functions /functions
RUN python3 -m pip install --upgrade pip
RUN pip3 install requests bs4
ENV fprocess="xargs bash /select_function.sh"
ENV write_debug="true"
EXPOSE 8080
HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1
CMD [ "fwatchdog" ]
