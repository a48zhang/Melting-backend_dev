FROM golang:1.19-alpine AS builder
ENV GOPROXY=https://goproxy.cn
ENV GIN_MODE=release
RUN mkdir -p src/log
ADD  . src
WORKDIR src
RUN go mod tidy && go build main
FROM alpine:latest
COPY --from=builder /go/src/ ./
CMD ["./main"]