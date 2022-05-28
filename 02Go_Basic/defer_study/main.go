package main

import "fmt"

func main() {
	//deferOrderDemo()
	//deferCopyDemo01(1,1) // 1 1
	//deferCopyDemo02(1,1) // 2 3
	deferAndReturnDemo01() //1 , 2
	deferAndReturnDemo02() //2 , 2
}

//1. defer 的调用顺序和栈一样，先进后出
func deferOrderDemo() {
	defer println("binshow")
	defer println("world! ")
	defer println("hello,")
}

//2. defer 调用函数参数的拷贝: defer将语句放入到栈中时，也会将相关的值拷贝同时入栈。
func deferCopyDemo01(num1 ,num2 int) {
	// 在对 num1 和 num2操作之前就复制到栈上了
	defer println("num2 = " , num2)
	defer println("num1 = " , num1)
	num1++
	num2 += 2
}

func deferCopyDemo02(num1 ,num2 int) {
	num1++
	num2 += 2
	defer println("num2 = " , num2)
	defer println("num1 = " , num1)
}


//3. defer and return 的返回时机
// (1) return 对返回变量赋值，如果是匿名返回值就先声明再赋值；
// (2) 执行 defer 函数；
// (3) return 携带返回值返回。

// 匿名返回值
func deferAndReturnDemo01() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2 value is ", i)
	}()

	defer func() {
		i++
		fmt.Println("defer1 in value is ", i)
	}()

	return i
}

// 有名返回值
func deferAndReturnDemo02() (j int) {
	defer func() {
		j++
		fmt.Println("defer2 in value", j)
	}()

	defer func() {
		j++
		fmt.Println("defer1 in value", j)
	}()

	return j
}

