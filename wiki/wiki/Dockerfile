FROM openfaas/classic-watchdog:0.13.4 as watchdog

FROM alpine:3.9

RUN apk add python3 git
RUN pip3 install requests 
RUN git clone https://github.com/ciankehoe/brickbot-wiki && cp brickbot-wiki/wiki.py .
COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

# Populate example here - i.e. "cat", "sha512sum" or "node index.js"
ENV fprocess="python3 wiki.py"
# Set to true to see request in function logs
ENV write_debug="false"

EXPOSE 8080

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]
