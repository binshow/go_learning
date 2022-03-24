package main

import "fmt"

func main() {
	// 声明一个变量 a
	var a int = 10

	// 声明一个指针，初始化为 a 的地址
	var p *int = &a

	fmt.Println("p = " , p) //0x1400012a008
	fmt.Println("address of a = " , &a) //0x1400012a008
	fmt.Println("*p = " , *p) // 10

	// 通过指针修改值
	*p = 20
	fmt.Println("a = " , a) // 20

	// 通过变量名修改值
	a = 30
	fmt.Println("a = " , a) //30
}

func test() {

	var a int = 10
	var b int = 20
	// a 在等号左边，认为是左值，代表变量所指向的内存空间
	// b 在等号右边，认为是右值，代表变量内存空间存储的数据值
	a = b
	fmt.Println(a)

}
