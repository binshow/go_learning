package gee

import (
	"fmt"
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()	// 表示执行了其他的handler
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}


func A() HandlerFunc {
	return func(c *Context) {
		fmt.Println("A enter .....")
		c.Next()
		fmt.Println("A end .....")
	}
}

func B() HandlerFunc {
	return func(c *Context) {
		fmt.Println("B enter .....")
		c.Next()
		fmt.Println("B end .....")
	}
}


func C() HandlerFunc {
	return func(c *Context) {
		fmt.Println("C enter .....")
		c.Next()
		fmt.Println("C end .....")
	}
}
