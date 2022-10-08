package main

import (
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-redis/redis/v7"
)

func main() {
	fmt.Println("golang连接redis")

	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "root&1234",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

}
