FROM golang:alpine  as builder
MAINTAINER Razil "woshilijinghua@gmail.com"
WORKDIR /kcos
RUN apk add  gcc g++ git
ADD . .
RUN go build



FROM alpine:3.11.6
RUN apk add  gcc g++ git
COPY --from=builder /kcos /kcos
WORKDIR /kcos
RUN mkdir /data
CMD ["/kcos"]
