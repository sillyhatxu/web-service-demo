FROM golang:1.14 AS builder

ENV PROJECT_NAME=github.com/sillyhatxu/web-service-demo
ENV MODULE_NAME=application-internal-api
COPY . $GOPATH/src/$PROJECT_NAME
WORKDIR $GOPATH/src/$PROJECT_NAME/$MODULE_NAME
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM xushikuan/alpine-build:1.0

ENV PROJECT_NAME=github.com/sillyhatxu/web-service-demo
ENV MODULE_NAME=application-internal-api
ENV BUILDER_SOURCE=/go/src/$PROJECT_NAME/$MODULE_NAME
ENV TIME_ZONE=Asia/Singapore
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

WORKDIR /app
COPY --from=builder $BUILDER_SOURCE/main .
ENTRYPOINT ./main