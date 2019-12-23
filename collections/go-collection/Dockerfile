FROM golang:1.12-alpine
RUN apk add --no-cache git make musl-dev go gcc curl bash
RUN echo "Pulling watchdog binary from Github." \
    && curl -sSL https://github.com/openfaas/faas/releases/download/0.9.6/fwatchdog > /usr/bin/fwatchdog \
    && chmod +x /usr/bin/fwatchdog \
    && cp /usr/bin/fwatchdog /
ADD select_function.sh /select_function.sh
ADD functions /functions
RUN go get github.com/skip2/go-qrcode
RUN go build /functions/pwned/pwned.go
RUN go build /functions/qr/qr.go
ENV fprocess="xargs bash /select_function.sh"
ENV write_debug="true"
EXPOSE 8080
HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1
CMD [ "fwatchdog" ]
