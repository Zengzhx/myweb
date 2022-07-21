FROM go-env:latest
ENV REDIS_HOST redis
ENV REDIS_PORT 6379
RUN mkdir -p /opt/myweb
WORKDIR /opt/myweb
COPY . .
RUN go build -o main
CMD ["./main"]