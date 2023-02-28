FROM golang:1.19-alpine
ENV GOPROXY=https://goproxy.cn GIN_MODE=release
RUN mkdir ~/build
ADD  . ~/build
WORKDIR ~/build
RUN go mod tidy & go build main -o ~/main & cd ~ & rm -rf ~/build & mkdir log
CMD ["~/main"]