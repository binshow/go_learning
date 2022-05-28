package main

import "fmt"

func main() {
/*
	// 声明了一个slice: 无需给定大小
	var ints []int

	// 声明了一个array: 需要指定大小
	var int [3]int

	fmt.Println("%T" , ints)
	fmt.Println("len(ints)=" , len(ints))
	fmt.Println("%T" , int)

	a := []string{"binshow" , "zkd" , "ustc"}
	a = append(a , "ahu")*/
	// 扩容如下：
	// 1. 扩容前容量 oldCap = 3 ， 添加一个元素最少要到 cap = 4
	//	  原容量翻倍 3 * 2 > 4 且 3 < 1024
	//	  所以预估容量直接翻倍 newCap = 3 * 2 = 6

	// 2. 预估容量 * 元素大小 ： 6 * 16 = 96 byte
	// 3. 匹配最近的内存规格: 8,16,32,48,64,80,96,112
	// 	  发现 96 是最接近的，所以 cap(a) = 96 / 16 = 6
	// 最后结果扩容后容量就是 6

	//SliceDemo()
	//SliceDemo01()
	SliceDemo03()
}

func SliceDemo() {
	// 创建 Slice 的方式:

	//1. 类型推导创建
	s1 := []int{1,2,3,4}
	fmt.Println("s1 = " , s1)

	//2. make
	s2 := make([]int , 5, 10)
	fmt.Println("len(s2) = " , len(s2) , "cap(s2) = " , cap(s2))

	s3 := make([]int , 7) // 未指定容量，则容量 = 长度
	fmt.Println("len(s3) = " , len(s3) , "cap(s3) = " , cap(s3))
}

func SliceDemo01() {
	// 声明了一个数组,要指定数组长度
	arr := []int{1,2,3,4,5}

	// 从数组截取Slice
	s := arr[1:3]
	fmt.Println("s = " , s)  //[2 , 3]
	fmt.Println("len(s) = " , len(s)) // 2
	fmt.Println("cap(s) = " , cap(s)) // 4

	s2 := arr[1:3:4] // 最后一个 4 指定了容量到数组的第4个元素
	fmt.Println("s2 = " , s2) //[2 , 3]
	fmt.Println("len(s2) = " , len(s2)) //2
	fmt.Println("cap(s2) = " , cap(s2)) //3
}

// SliceDemo02 append 函数测试
func SliceDemo02() {
	// 调用append往Slice追加元素时，会自动扩容
	s1 := []int{1,2,3,4}
	fmt.Println("s1 = " , s1)
	s1 = append(s1 , 5)
	fmt.Println("s1 = " , s1)
	s1 = append(s1 , 6)
	fmt.Println("s1 = " , s1)
}


// SliceDemo03 copy 函数测试
func SliceDemo03() {

	data := []int{0,1,2,3,4,5,6,7,8,9}

	s1 := data[8:]
	s2 := data[0:5]
	fmt.Printf("s1 = %v , s2 = %v\n" , s1 , s2) //s1 = [8 9] , s2 = [0 1 2 3 4]

	copy(s1 , s2) // 将 s2 切片拷贝到 s1 的位置
	fmt.Printf("s1 = %v , s2 = %v\n" , s1 , s2) //s1 = [0 1] , s2 = [0 1 2 3 4]
}
