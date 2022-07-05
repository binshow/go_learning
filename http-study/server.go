package main

import "net/http"

// -------------------------------------------
// @file          : server.go
// @author        : binshow
// @time          : 2022/6/22 11:01 PM
// @description   : 自定义简单http server 实现
// -------------------------------------------

type Server interface {

	//Route(pattern string , handleFunc http.HandlerFunc)

	//Route(pattern string , handleFunc func(ctx *Context)) // 在这里需要改方法签名了,支持 context 在框架层创建


	Route(method string , pattern string , handleFunc func(ctx *Context)) // 支持restful api，所以需要传入 method

	Start(address string) error
}

// 根据http
type sdkHttpServer struct {
	Name string
	handler *HandlerBasedOnMap
}

// Route 注册路由
func (s *sdkHttpServer) Route( method , pattern string, handleFunc func(ctx *Context)) {
	// 希望在这里创建 context , 可以通过闭包的方式如下创建
	//http.HandleFunc(pattern , func(writer http.ResponseWriter, request *http.Request) {
	//	ctx := NewContext(writer , request)
	//	handleFunc(ctx)
	//})
	//http.Handle(pattern , &HandlerBasedOnMap{})

	key := s.handler.key(method , pattern)
	s.handler.handlers[key] = handleFunc

}

func (s *sdkHttpServer) Start(address string) error {
	http.Handle("/" , s.handler) // 只注册一遍，在启动的时候注册
	return http.ListenAndServe(address , nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}