package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go_learning/protobuf_入门/demo01/example"
	"os"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/7/6 10:17 AM
// @description   :
// -------------------------------------------

func main() {

	fmt.Println("Hello World")


	orderTest := &example.Order{
		OrderId:   1,
		Num:       20,
		Timestamp: "20180415",
	}


	//序列化
	msgDataEncoding, err := proto.Marshal(orderTest)
	if err != nil {
		panic(err.Error())
		return
	}

	//反序列化：
	msgEntity := example.Order{}
	err = proto.Unmarshal(msgDataEncoding, &msgEntity)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
		return
	}

	fmt.Println(msgEntity.GetOrderId())
	fmt.Println(msgEntity.GetTimestamp())
	fmt.Println(msgEntity.GetNum())

}