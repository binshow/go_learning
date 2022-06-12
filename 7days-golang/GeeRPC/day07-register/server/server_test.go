package server

import (
	"fmt"
	"github.com/mailru/easyjson/buffer"
	"log"
	"reflect"
	"strings"
	"testing"
)

// -------------------------------------------
// @file          : server_test.go
// @author        : binshow
// @time          : 2022/6/9 10:41 PM
// @description   : 将结构体映射为服务
// -------------------------------------------


/*

对 net/rpc 而言，一个函数需要能够被远程调用，需要满足如下五个条件：
the method’s type is exported. – 方法所属类型是导出的。
the method is exported. – 方式是导出的。
the method has two arguments, both exported (or builtin) types. – 两个入参，均为导出或内置类型。
the method’s second argument is a pointer. – 第二个入参必须是一个指针。
the method has return type error. – 返回值为 error 类型

如下：
func (t *T) MethodName(argType T1, replyType *T2) error

*/


//通过反射来获取某个结构体的所有方法，并且通过方法获取到该方法所有的参数类型和返回值
func TestReflectGetStruct(t *testing.T) {
	//var wg sync.WaitGroup
	var buf buffer.Buffer
	typ := reflect.TypeOf(&buf)
	fmt.Printf("typ = %v\n" , typ) // yp = *sync.WaitGroup

	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		argv := make([]string , 0 , method.Type.NumIn())
		returns := make([]string , 0 , method.Type.NumOut())
		// j 从 1 开始，第 0 个入参是 wg 自己。
		for j := 1; j < method.Type.NumIn(); j++ {
			argv = append(argv , method.Type.In(j).Name())
		}

		for j := 1; j < method.Type.NumOut(); j++ {
			returns = append(returns , method.Type.In(j).Name())
		}

		log.Printf("func (w *%s) %s(%s) %s",
			typ.Elem().Name(),
			method.Name,
			strings.Join(argv, ","),
			strings.Join(returns, ","))
	}

}

type Foo int
type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

// it's not a exported Method
func (f Foo) sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func _assert(condition bool, msg string, v ...interface{}) {
	if !condition {
		panic(fmt.Sprintf("assertion failed: "+msg, v...))
	}
}

func TestNewService(t *testing.T) {
	var foo Foo
	s := newService(&foo)
	fmt.Println(s.name) // Foo
	//s.rcvr = <*day02_client.Foo Value> ,  s.method = map[Sum:0x14000140080] , s.typ = *day02_client.Foo
	fmt.Printf("s.rcvr = %v ,  s.method = %v , s.typ = %v \n" , s.rcvr.String() , s.method , s.typ.String() )
	_assert(len(s.method) == 1, "wrong service Method, expect 1, but got %d", len(s.method))
	mType := s.method["Sum"]
	_assert(mType != nil, "wrong Method, Sum shouldn't nil")
}

func TestMethodType_Call(t *testing.T) {
	var foo Foo
	s := newService(&foo)
	mType := s.method["Sum"]

	argv := mType.newArgv()
	replyv := mType.newReplyv()
	argv.Set(reflect.ValueOf(Args{Num1: 1, Num2: 3}))
	err := s.call(mType, argv, replyv)
	_assert(err == nil && *replyv.Interface().(*int) == 4 && mType.NumCalls() == 1, "failed to call Foo.Sum")
}
