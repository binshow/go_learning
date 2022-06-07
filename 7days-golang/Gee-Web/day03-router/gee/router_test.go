package gee

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

/**
实现动态路由
		      /
	/hello    /hi     /assets
/:name /b	  /:name  /*filepath
	   /c
*/

func newTestRouter() *router {
	r := newRouter() //{}
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)

	bytes, err := json.Marshal(r.roots)
	if err != nil {
		fmt.Printf("err = %v", err)
	}
	fmt.Println(string(bytes))

	bytes2, err := json.Marshal(r.handlers)
	if err != nil {
		fmt.Printf("err = %v" , err) //err = json: unsupported type: gee.HandlerFunc
	}
	fmt.Println(string(bytes2))
	return r
}

// 测试路径解析
func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"}) // 只允许一个 *
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

// 测试获取路由
func TestGetRoute(t *testing.T) {
	r := newTestRouter()

	//n, ps := r.getRoute("GET", "/hello/geektutu")
	n, ps := r.getRoute("GET", "/hello/b")
	n.String()

	//if n == nil {
	//	t.Fatal("nil shouldn't be returned")
	//}
	//
	//if n.pattern != "/hello/:name" {
	//	t.Fatal("should match /hello/:name")
	//}
	//
	//if ps["name"] != "geektutu" {
	//	t.Fatal("name should be equal to 'geektutu'")
	//}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])

}

func TestGetRoute2(t *testing.T) {
	r := newTestRouter()
	n1, ps1 := r.getRoute("GET", "/assets/file1.txt")
	ok1 := n1.pattern == "/assets/*filepath" && ps1["filepath"] == "file1.txt"
	if !ok1 {
		t.Fatal("pattern shoule be /assets/*filepath & filepath shoule be file1.txt")
	}

	n2, ps2 := r.getRoute("GET", "/assets/css/test.css")
	ok2 := n2.pattern == "/assets/*filepath" && ps2["filepath"] == "css/test.css"
	if !ok2 {
		t.Fatal("pattern shoule be /assets/*filepath & filepath shoule be css/test.css")
	}

}

func TestGetRoutes(t *testing.T) {
	r := newTestRouter()
	nodes := r.getRoutes("GET")
	for i, n := range nodes {
		fmt.Println(i+1, n)
	}

	if len(nodes) != 5 {
		t.Fatal("the number of routes shoule be 4")
	}
}
