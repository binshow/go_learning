package main

import "fmt"

func main() {

	// 1. 用make函数初始化切片时，如果不指明其容量，那么它就会和长度一致
	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1)) // 5
	fmt.Printf("The capacity of s1: %d\n", cap(s1)) //5
	fmt.Printf("The value of s1: %d\n", s1)

	//2. 切片的容量实际上代表了它的底层数组的长度
	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2)) // 5
	fmt.Printf("The capacity of s2: %d\n", cap(s2)) // 8
	fmt.Printf("The value of s2: %d\n", s2)
	fmt.Println()

	// 示例2。
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]
	fmt.Printf("The length of s4: %d\n", len(s4))
	fmt.Printf("The capacity of s4: %d\n", cap(s4))
	fmt.Printf("The value of s4: %d\n", s4)
	fmt.Println()

	// 示例3。
	s5 := s4[:cap(s4)]
	fmt.Printf("The length of s5: %d\n", len(s5))
	fmt.Printf("The capacity of s5: %d\n", cap(s5))
	fmt.Printf("The value of s5: %d\n", s5)

}
