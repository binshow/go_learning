package main

import (
	"net/http"
)

// -------------------------------------------
// @file          : map_based_handler.go
// @author        : binshow
// @time          : 2022/6/22 11:41 PM
// @description   : 基于map实现handler
// -------------------------------------------

type HandlerBasedOnMap struct {
	handlers map[string]func(ctx *Context) // key = method + url , val = 对应的handler函数
}


func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	key := h.key(request.Method , request.URL.Path)
	if handler , ok := h.handlers[key] ; ok {
		// 找到了对应的路由处理
		handler(NewContext(writer , request))
	}else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Not Found"))
	}

}

func (h *HandlerBasedOnMap) key(method string , path string) string {
	return method + "#" + path // Path 是肯定有的！！！
}

