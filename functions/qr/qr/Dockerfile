FROM openfaas/classic-watchdog:0.18.0 as watchdog

FROM golang:1.12-alpine

RUN apk add git

COPY ./qr.go /qr.go

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

RUN go get github.com/skip2/go-qrcode

RUN go build /qr.go
# Populate example here - i.e. "cat", "sha512sum" or "node index.js"
ENV fprocess="xargs /go/qr"
# Set to true to see request in function logs
ENV write_debug="false"

EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]
