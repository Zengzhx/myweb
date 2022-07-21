package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"net/http"
	"os"
)

var rd *redis.Client

func main() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	if redisHost == "" {
		redisHost = "redis"
	}
	if redisPort == "" {
		redisPort = "6379"
	}
	addr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	fmt.Printf("redis连接配置:%v\n", addr)
	rd = redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    addr,
	})
	cmd := rd.Do("ping")
	if cmd.Err() != nil {
		fmt.Fprintf(os.Stderr, "redis 连接失败 - %v\n", cmd.Err())
		return
	}
	http.HandleFunc("/hello", handlerHello)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "web 服务开启失败 - %v\n", err)
		return
	}
}

func handlerHello(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("[%v] - I'm Recive \n", req.Method)
	iniCmd := rd.Incr("views")
	if iniCmd.Err() != nil {
		fmt.Fprintf(os.Stderr, "访问量统计失败- %v\n", iniCmd.Err())
	}
	v := fmt.Sprintf("hello, views:%v", iniCmd.Val())
	res.Write([]byte(v))
	res.WriteHeader(http.StatusOK)
}
