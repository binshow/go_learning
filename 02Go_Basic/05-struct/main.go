package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Book struct {
	Title string
	Pages int
	Indexes map[string]int // 索引
}

type T struct {
	b byte
	i int64
	u uint16
}


type S struct {
	b byte
	u uint16
	i int64
}

func main() {
	var t T
	fmt.Println(unsafe.Sizeof(t)) // 24
	var s S
	fmt.Println(unsafe.Sizeof(s)) // 16


	//CompareStruct01() // a1 == a2
	//CompareStruct02() // a1 != a2
	////CompareStruct03() //
	////CompareStruct04() //
}

func CompareStruct01() {
	type A struct {
		name 	string
		address string
	}
	a1 := A{"jian", "ustc"}
	a2 := A{"jian", "ustc"}
	if a1 == a2 {
		fmt.Println("a1 == a2")
		return
	}
	fmt.Println("a1 != a2")
}

func CompareStruct02() {
	type A struct {
		name 	string
		address *string
	}
	a1 := A{"binshow", new(string)}
	a2 := A{"binshow", new(string)}
	if a1 == a2 {
		fmt.Println("a1 == a2")
		return
	}
	fmt.Println("a1 != a2")
}

//func CompareStruct03() {
//	type A struct {
//		name 	string
//		address []string
//	}
//	a1 := A{"binshow", []string{"a","b","c"}}
//	a2 := A{"binshow", []string{"a","b","c"}}
//	if a1 == a2 {  //Invalid operation: a1 == a2 (the operator == is not defined on A)
//		fmt.Println("a1 == a2")
//		return
//	}
//	fmt.Println("a1 != a2")
//}
//
//// 不同结构体相同值比较
//func CompareStruct04() {
//	type A struct {
//		name 	string
//	}
//	type B struct {
//		name 	string
//	}
//
//	a1 := A{"binshow"}
//	a2 := B{"binshow"}
//	// 同样编译报错，虽然可以强制类型转换
//	if a1 == a2 {  //Invalid operation: a1 == a2 (the operator == is not defined on A)
//		fmt.Println("a1 == a2")
//		return
//	}
//	if a1 == A(a2) {
//		fmt.Println("a1 == a2")
//		return
//	}
//	fmt.Println("a1 != a2")
//}

func CompareStruct05() {
	type A struct {
		name 	string
		address *string
	}
	a1 := A{"binshow", new(string)}
	a2 := A{"binshow", new(string)}
	if reflect.DeepEqual(a1, a2) {
		fmt.Println("a1 == a2")
		return
	}
	fmt.Println("a1 != a2")
}






