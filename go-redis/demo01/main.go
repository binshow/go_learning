package main

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/7/19 11:04 PM
// @description   : go 操作redis
// -------------------------------------------

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)


var key string = "zhanqu_1"
var rdb *redis.Client

func initRedisClient(){
	rdb = redis.NewClient(&redis.Options{
		Username: "",
		Password: "",
		Addr: "localhost:6379",
		DB: 0,
		PoolSize: 5,
	})
}



type AnswerRateBound struct {
	LowBound  float64 `json:"low_bound"`
	HighBound float64 `json:"high_bound"`
	IsNotify  int32   `json:"is_notify"`
}



func main() {
	ctx := context.Background()
	//// ping一下检查是否连通
	//pingResult, err := rdb.Ping(ctx).Result()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//// PONG
	//fmt.Println(pingResult)
	initRedisClient()

	bound := &AnswerRateBound{
		LowBound:  1,
		HighBound: 10,
		IsNotify:  0,
	}
	bytes, _ := json.Marshal(bound)
	res := rdb.Set(ctx, key, bytes, 10*60)
	fmt.Println(res.String())

	cache := rdb.Get(ctx, key)
	fmt.Println(cache.Err())
}

//1. 第一个请求进来发现需要调价
func firstReq(i int) {
	ctx := context.Background()
	// 调价了
	// 读redis 看是否已经发过通知了
	cache := rdb.Get(ctx, key)
	if cache.Err() != nil {
		// 说明redis中还没有，需要写入
		bound := &AnswerRateBound{
			LowBound:  5,
			HighBound: 10,
			IsNotify:  0,
		}
		bytes, _ := json.Marshal(bound)
		_ = rdb.Set(ctx, key, bytes, 10*60)
		fmt.Println("第" , i  , "个发送通知了。。。。")
	}else {
		bound := &AnswerRateBound{}
		bytes, _ := cache.Bytes()
		_ = json.Unmarshal(bytes, bound)
		if bound.IsNotify == 1 {

		}
	}
}