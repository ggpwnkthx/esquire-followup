FROM golang:1-alpine as builder
RUN apk add --no-cache --virtual .build-deps gcc libc-dev git
COPY ./ /go/src/app
WORKDIR /go/src/app
RUN go build -o /bin/go-app

FROM alpine:3
COPY --from=builder /bin/go-app /bin/go-app
EXPOSE 8080
ENTRYPOINT [ "/bin/go-app" ]
