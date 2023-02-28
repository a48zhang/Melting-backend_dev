FROM golang:alpine
ENV GOPROXY=https://goproxy.cn
ENV GIN_MODE=release
RUN mkdir build/
ADD  . /build
WORKDIR /build
RUN go mod tidy & go build main & cp ./main /root/main
WORKDIR /root
RUN rm -rf ~/build & mkdir log
CMD ["/root/main"]