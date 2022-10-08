package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // redis 密码
		DB:       0,  // 默认数据库，默认为0
	})

	err := client.Set("key", "value", 0).Err()

	// 检测是否连接成功
	if err != nil {
		panic(err)
	}

	// 根据key查询缓存，通过Result返回两个值
	val, err := client.Get("key").Result()

	// 检查错误
	if err != nil {
		panic(err)
	}

	fmt.Println("value", val)
}
