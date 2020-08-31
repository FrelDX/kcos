FROM golang AS build
ENV GOPROXY https://goproxy.cn
COPY . /kcos
RUN CGO_ENABLED=1 GOOS=linux
RUN cd /kcos &&  go build

WORKDIR /
COPY key/id_rsa /key/id_rsa
CMD ["/kcos/kcos"]
