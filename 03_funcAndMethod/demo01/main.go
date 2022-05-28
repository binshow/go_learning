package main

import "fmt"


// 这里声明的a 和 b 就是形参
func Add(a int, b int) (res int) {
	res = a + b
	return
}

func main() {
	// 这里传入的 1 和 2 就是实参
	res := Add(1, 2)
	fmt.Println(res)
}
