package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// -------------------------------------------
// @file          : main.go
// @author        : binshow
// @time          : 2022/6/22 11:01 PM
// @description   :
// -------------------------------------------

type signUpReq struct {
	Email  string `json:"email"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type commonResponse struct {
	BizCode int
	Msg     string
	Data    interface{}
}


func main() {
	server := NewHttpServer("test-server")
	//server.Route("/" , home)
	//server.Route("/signUp" , signUpWithoutContext)

	// server.Route("/signUp2" , signUp2) context 需要在 Route 的实现中创建
	server.Start(":8080")
}

// signUp0 实现注册用户
func signUpWithoutContext(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}
	// 读请求中的body数据
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w , "read body failed : %v" , err)
		return
	}

	err = json.Unmarshal(body ,req)
	if err != nil {
		fmt.Fprintf(w , "deserialized failed : %v" , err)
		return
	}


	// 业务逻辑 ...




	// 返回一个通用的resp
	resp := &commonResponse{
		Data: 123,
	}
	respJson, err := json.Marshal(resp)
	if err != nil {

	}
	fmt.Fprintf(w , string(respJson))

}

// signUp1 实现注册用户，封装了 context
func signUp(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{R: r,W: w}

	req := &signUpReq{}
	err := ctx.ReadJson(req)
	if err != nil {
		fmt.Fprintf(w , "deserialized failed : %v" , err)
		return
	}

	// 业务逻辑 ...


	// 返回一个通用的resp
	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.WriterJson(http.StatusOK , resp)
	if err != nil {
		// 这里只能写日志了
		fmt.Printf("write resp failed : %v" , err)
		return
	}


}


// signUp2 不想将这个context的创建交给context，而是交给框架层去创建, 那么框架层怎么创建context呢？？？
func signUp2(ctx *Context) {
	req := &signUpReq{}
	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}

	// 业务逻辑 ...


	// 返回一个通用的resp
	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.WriterJson(http.StatusOK , resp)
	if err != nil {
		// 这里只能写日志了
		fmt.Printf("write resp failed : %v" , err)
		return
	}


}

func home(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer , "this is home")
}



