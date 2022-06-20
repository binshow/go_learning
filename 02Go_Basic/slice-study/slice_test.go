package slice_study

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

// -------------------------------------------
// @file          : slice_test.go
// @author        : binshow
// @time          : 2022/6/20 9:40 AM
// @description   : 关于 slice 的常见面试题
// -------------------------------------------


// 基础知识： slice 在go中由一下数据结构表示：
//type SliceHeader struct {
//	Data uintptr
//	Len  int
//	Cap  int
//}


//1. nil 切片 和 空 切片的区别
func TestQuest1(t *testing.T) {
	var s1  []int			// nil 切片
	s2 := make([]int , 0)	// 空  切片
	s4 := make([]int , 0)
	fmt.Printf("s1 pointer:%+v, s2 pointer:%+v, s4 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&s1)),*(*reflect.SliceHeader)(unsafe.Pointer(&s2)),*(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s1))).Data==(*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data)
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&s2))).Data==(*(*reflect.SliceHeader)(unsafe.Pointer(&s4))).Data)

	// 答案：
	// 1. nil切片和空切片指向的地址不一样。nil空切片引用数组指针地址为0（无指向任何实际地址）
	// 2. 空切片的引用数组指针地址是有的，且固定为一个值
}