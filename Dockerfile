FROM golang AS build
ENV GOPROXY https://goproxy.cn
COPY . /kcos
RUN CGO_ENABLED=1 GOOS=linux
RUN cd /kcos &&  go build

WORKDIR /kcos
COPY key/id_rsa /kcos/key/id_rsa
CMD ["./kcos"]
