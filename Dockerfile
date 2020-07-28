FROM golang:alpine  as builder
MAINTAINER Razil "woshilijinghua@gmail.com"
WORKDIR /kcos
RUN apk add  gcc g++ git
ADD . .
RUN CGO_ENABLED=1 GOOS=linux
RUN go build



FROM alpine:3.5
RUN apk add  gcc g++ git
COPY --from=builder /kcos /kcos
WORKDIR /kcos
RUN mkdir /data
CMD ["/kcos"]
