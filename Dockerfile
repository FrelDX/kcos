FROM golang:alpine  as builder
MAINTAINER Razil "woshilijinghua@gmail.com"
WORKDIR /kcos
ADD . .
RUN go build



FROM alpine:3.11.6
WORKDIR /kcos
COPY --from=builder /kcos/kcos /kcos
RUN mkdir /data
CMD ["/kcos"]
