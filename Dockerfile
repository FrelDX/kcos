FROM alpine:edge AS build
RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.9/main" > /etc/apk/repositories \
    && echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.9/community" >> /etc/apk/repositories
RUN apk update
RUN apk upgrade
RUN apk add  go gcc g++
WORKDIR /
ENV GOPATH /
ENV GOPROXY https://goproxy.cn
COPY . /kcos
RUN CGO_ENABLED=1 GOOS=linux
RUN cd /kcos &&  go build


FROM alpine:3.5
WORKDIR /
COPY --from=build /kcos/kcos /kcos
RUN mkdir /data
CMD ["/kcos"]

