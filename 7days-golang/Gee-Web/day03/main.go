package main

import (
	gee2 "go_learning/7days-golang/Gee-Web/day03/gee"
	"net/http"
)

func main() {
	r := gee2.New()

	//r.GET("/", func(c *gee.Context) {
	//	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	//})

	r.GET("/hello/test", func(c *gee2.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gee2.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee2.Context) {
		c.JSON(http.StatusOK, gee2.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}