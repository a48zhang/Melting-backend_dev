FROM golang:1.19-alpine
ENV GOPROXY=https://goproxy.cn GIN_MODE=release
RUN mkdir melting/build -p ; mkdir melting/bin/log -p
ADD  . /melting
WORKDIR /melting/build
RUN go mod tidy ; go build main ;cp ./main /melting/bin
WORKDIR melting/bin
CMD ["./main"]
