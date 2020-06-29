FROM golang AS build
ENV GOPROXY https://goproxy.cn
COPY . /kube-console-on-ssh
RUN CGO_ENABLED=1 GOOS=linux
RUN cd /kube-console-on-ssh &&  go build




FROM alpine:3.5
WORKDIR /
COPY --from=build /kube-console-on-ssh/kube-console-on-ssh /kube-console-on-ssh
COPY key/id_rsa /key/id_rsa
CMD ["/kube-console-on-ssh"]
