package main

import (
	"fmt"
	"sync"
)

// 使用 Once 来实现单例模式


// 首字母小写
type singleton struct {}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	// once.Do 可以保证这个函数在整个程序运行过程中都只运行一次
	once.Do(func() {
		instance = new(singleton)  // 初始化
	})
	return instance
}



func main()  {

}

// 检查下面三个函数的输出结果
func panicDo()  {
	once := &sync.Once{}
	defer func() {
		if err := recover();err != nil{
			once.Do(func() {
				fmt.Println("run in recover")
			})
		}
	}()
	once.Do(func() {
		panic("panic i=0")
	})

}

func nestedDo()  {
	once := &sync.Once{}
	once.Do(func() {
		once.Do(func() {
			fmt.Println("test nestedDo")
		})
	})
}

func nestedDo2()  {
	once1 := &sync.Once{}
	once2 := &sync.Once{}
	once1.Do(func() {
		once2.Do(func() {
			fmt.Println("test nestedDo")
		})
	})
}


