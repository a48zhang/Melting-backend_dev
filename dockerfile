FROM golang:alpine
ENV GOPROXY=https://goproxy.cn
ENV GIN_MODE=release
RUN mkdir build/
ADD  . /build
WORKDIR /build
RUN go mod tidy
RUN go build main
RUN cp ./main ~/main
WORKDIR ~
RUN rm -rf ~/build
RUN mkdir log
CMD ["~/main"]