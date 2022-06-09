package main

/*
	中间件：middleware 是框架提供的一个插口，允许用户自己定义功能，嵌入到框架中去
	需要考虑的地方：
	1. 插入点在哪？ 太底层逻辑就比较复杂，太上层比如用户自己定义一个函数去执行
	2. 中间件输入是啥？ 输入参数的多少代表了中间件的扩展能力

	最后的设计：middleware 应该作用于 routeGroup 路由分组上。
	当收到请求后，匹配路由，该请求的所有信息都保存在 context 中，所以我们也应该将 middleware 保存到context中去依次调用
*/

import (
	gee2 "go_learning/7days-golang/Gee-Web/day05-middleware/gee"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gee2.HandlerFunc {
	return func(c *gee2.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gee2.New()

	//r.Use(gee2.Logger()) // global midlleware

	r.Use(gee2.A() , gee2.B() , gee2.C())  // 1.测试中间件注册和执行的顺序  A B C 依次 enter , C B A 依次 end

	r.GET("/", func(c *gee2.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		//v2.GET("/hello/:name", func(c *gee2.Context) {
		//	// expect /hello/geektutu
		//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		//})
	}

	r.Run(":9998")
}
