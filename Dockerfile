FROM golang:1.6
RUN mkdir -p /opt/myweb
RUN go build -o main
ENV REDIS_HOST redis
ENV REDIS_PORT 6379
WORKDIR /opt/myweb
CMD ["./main"]