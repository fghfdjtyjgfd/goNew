FROM golang:1.20.5-alpine3.18 AS builder

ENV GOOS=linux
ENV GOARCH=amd64
RUN GOCACHE=OFF
ENV GO111MODULE=on
ENV TZ=Asia/Bangkok

RUN apk update && apk upgrade && \
    apk add --no-cache ca-certificates git wget
RUN apk add build-base

RUN mkdir /api
WORKDIR /api
ADD . /api

RUN go build -o goapp main.go
RUN apk add --no-cache tzdata


FROM alpine:3.15

COPY --from=builder /api/goapp .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /api/configs/ ./configs/

ENTRYPOINT ["/goapp"]