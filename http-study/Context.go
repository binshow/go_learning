package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// -------------------------------------------
// @file          : Context.go
// @author        : binshow
// @time          : 2022/6/22 11:17 PM
// @description   :	Context 包装
// -------------------------------------------

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{W: w, R: r}
}

func (c *Context) ReadJson(req interface{}) error {
	// 读了 request中的body 并反序列化
	r := c.R
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body ,req)
	if err != nil {
		return err
	}
	return nil
}

//WriterJson 核心方法
func (c *Context) WriterJson(code int , resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	_, err = c.W.Write(respJson)
	return err
}

//====== 下面三个是通过核心方法实现的辅助方法

// OkJson 直接返回成功的响应码 ， 辅助方法，通过核心方法来实现！！！
func (c *Context) OkJson( resp interface{}) error {
	return c.WriterJson(http.StatusOK , resp)
}

func (c *Context) SystemErrorJson( resp interface{}) error {
	return c.WriterJson(http.StatusInternalServerError , resp)
}

func (c *Context) BadRequestJson( resp interface{}) error {
	return c.WriterJson(http.StatusNotFound , resp)
}

