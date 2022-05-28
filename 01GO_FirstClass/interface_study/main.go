package main

import "fmt"



func main() {
	printNonEmptyInterface()
}


func printNilInterface() {
	// nil接口变量
	var i interface{} // 空接口类型
	var err error     // 非空接口类型

	// 无论是空接口类型还是非空接口类型变量，一旦变量值为 nil，
	// 那么它们内部表示均为(0x0,0x0)，也就是类型信息、数据值信息均为空
	println(i) 	 // (0x0,0x0)
	println(err) // (0x0,0x0)
	println("i = nil:", i == nil) 	// true
	println("err = nil:", err == nil) // true
	println("i = err:", i == err)		// true
}


func printEmptyInterface() {
	var eif1 interface{} // 空接口类型
	var eif2 interface{} // 空接口类型
	var n, m int = 17, 18

	eif1 = n
	eif2 = m

	println("eif1:", eif1) // (0x100867f80,0x1400005cf60)
	println("eif2:", eif2) // (0x100867f80,0x1400005cf58)
	println("eif1 = eif2:", eif1 == eif2) // false

	eif2 = 17
	println("eif1:", eif1)  // (0x100867f80,0x1400005cf60)
	println("eif2:", eif2)  // (0x100867f80,0x100863178)
	println("eif1 = eif2:", eif1 == eif2) // true

	eif2 = int64(17)
	println("eif1:", eif1) // (0x100867f80,0x1400005cf60)
	println("eif2:", eif2) // (0x100868040,0x100863178)
	println("eif1 = eif2:", eif1 == eif2) // false
}


type T int

func (t T) Error() string {
	return "bad error"
}

// 和空接口类型变量一样，只有 tab 和 data 指的数据内容一致的情况下，
// 两个非空接口类型变量之间才能划等号
func printNonEmptyInterface() {
	var err1 error // 非空接口类型
	var err2 error // 非空接口类型
	err1 = (*T)(nil)
	println("err1:", err1) // err1: (0x104fc4300,0x0)
	println("err1 = nil:", err1 == nil) // false

	err1 = T(5)
	err2 = T(6)
	println("err1:", err1) // err1: (0x104fc4360,0x104fa9840)
	println("err2:", err2) // err2: (0x104fc4360,0x104fa9848)
	println("err1 = err2:", err1 == err2) // false

	err2 = fmt.Errorf("%d\n", 5)
	println("err1:", err1) // err1: (0x104fc4360,0x104fa9840)
	println("err2:", err2) // err2: (0x104fc4280,0x14000104210)
	println("err1 = err2:", err1 == err2) // false
}
