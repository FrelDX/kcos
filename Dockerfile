FROM alpine:edge AS build
RUN echo "https://mirrors.aliyun.com/alpine/v3.9/main" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.9/community" >> /etc/apk/repositories
RUN apk update
RUN apk upgrade
RUN apk add  go gcc g++ git
COPY ./*  /kcos/
RUN CGO_ENABLED=1 GOOS=linux	RUN CGO_ENABLED=1 GOOS=linux
RUN cd /kcos/kcos &&  go build	WORKDIR /kcos
COPY key/id_rsa /key/id_rsa	RUN ls -l && pwd
WORKDIR /kube-console-on-ssh/	RUN go build
CMD ["./kube-console-on-ssh"]	



FROM alpine:3.5

COPY --from=build /kcos/kcos /kcos
WORKDIR /kcos
RUN mkdir /data
CMD ["/kcos"]
