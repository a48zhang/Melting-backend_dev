FROM golang:1.19-alpine
ENV GOPROXY=https://goproxy.cn GIN_MODE=release
RUN mkdir melting
ADD  . /melting
WORKDIR /melting
RUN mkdir log ; go mod tidy ; go build main
CMD ["./main"]
