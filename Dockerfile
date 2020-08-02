FROM golang:1.13-alpine AS builder
ENV PACKAGE=github.com/drugs-4-3/ingrid_task
ENV GOOS=linux
ENV LISTEN=0.0.0.0:8080
ENV CGO_ENABLED=0
WORKDIR /go/src/${PACKAGE}
RUN apk update && apk add curl && apk add git && apk add bash && apk add --no-cache git tzdata bash
RUN go get -u github.com/golang/dep/cmd/dep
COPY . .
RUN dep init || true
RUN dep ensure
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build  -o /go/bin/ingrid_task

FROM alpine:3.9
COPY --from=builder /go/bin/ingrid_task /go/bin/ingrid_task
CMD ["./go/bin/ingrid_task"]
EXPOSE 8080
