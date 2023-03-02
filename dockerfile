FROM golang:1.19-alpine AS builder
ENV GOPROXY=https://goproxy.cn
RUN mkdir -p src/log
ADD  . src
WORKDIR src
RUN go mod tidy && go build main
FROM alpine:latest
ENV GIN_MODE=release
COPY --from=builder /go/src/ ./
CMD ["./main"]