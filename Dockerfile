FROM golang AS build
COPY . /kube-console-on-ssh
RUN CGO_ENABLED=1 GOOS=linux
RUN cd /kube-console-on-ssh &&  go build
COPY key/id_rsa /key/id_rsa
WORKDIR /kube-console-on-ssh/
CMD ["./kube-console-on-ssh"]
