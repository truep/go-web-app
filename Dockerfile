FROM golang:latest as builder

WORKDIR /opt/code
ADD ./ /opt/code

RUN make

FROM debian:latest
ENV APP_NAME web-app-truep


WORKDIR /opt/app

COPY --from=builder /opt/code/bin/${APP_NAME} ./
COPY --from=builder /opt/code/config.yml ./

EXPOSE 8080

ENTRYPOINT ["./web-app-truep"]