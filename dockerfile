FROM golang:1.19-alpine
ENV GOPROXY=https://goproxy.cn GIN_MODE=release
RUN mkdir melting
ADD  . /melting
WORKDIR /melting
RUN mkdir log
RUN go mod tidy
RUN go build main
CMD ["./main"]
