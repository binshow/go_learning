package chapter_3_data_structer

import (
	"fmt"
	"testing"
)

// -------------------------------------------
// @file          : slice_test.go
// @author        : binshow
// @time          : 2022/7/6 1:16 PM
// @description   :	切片Slice的设计与实现
// -------------------------------------------

// 编译期间的切片是 cmd/compile/internal/types.Slice 类型的
func TestNewSlice(t *testing.T) {
	//1. 使用数组下标
	arr := [3]int{1,2,3}
	slice1 := arr[0:1]
	fmt.Println("slice1 = " , slice1)

	//2. 使用字面量直接声明
	slice2 := []int{1,2,3,4}
	fmt.Println("slice2 = " , slice2)

	//3. 使用new关键字
	slice3 := make([]int , 0 , 3)
	fmt.Println("slice3 = " , slice3)
}


//在运行时切片可以由如下的 reflect.SliceHeader 结构体表示

func TestRunTimeSlice(t *testing.T) {
	//reflect.SliceHeader{}


}