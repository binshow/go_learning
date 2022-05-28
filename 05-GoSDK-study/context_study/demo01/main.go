package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-uuid"
	"net/http"
)

func main() {
	// curl -v http://localhost:8080/welcome
	ContextWithValueDemo()
}


// ContextWithValueDemo 测试 WithValue() func
// 对每个进来的请求request在header中都注入 msgID
func ContextWithValueDemo() {
	helloWorldHandler := http.HandlerFunc(HelloWorld)
	http.Handle("/welcome" , injectMsgId(helloWorldHandler))
	if err := http.ListenAndServe(":8080", nil); err != nil{
		fmt.Printf("ListenAndServe err: %v\n" , err)
	}
}


func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msgID := ""
	if m := r.Context().Value("msgId"); m != nil{
		if value , ok := m.(string) ; ok{
			msgID = value
		}
	}
	w.Header().Add("msgId" , msgID)
	w.Write([]byte("Hello , world!"))
}


func injectMsgId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			msgID , _ := uuid.GenerateUUID()
			ctx := context.WithValue(r.Context() , "msgId" , msgID)
			req := r.WithContext(ctx)
			next.ServeHTTP(w , req)
	})
}
