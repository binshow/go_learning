package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var s string = "hello"
	// 通过 unsafe.Pointer 通用指针转型能力，找到底层数组的地址
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("0x%x\n" , hdr.Data) //0x102a51797
	fmt.Printf("len(s) = %v\n" , hdr.Len) //5
	p := (*[5]byte)(unsafe.Pointer(hdr.Data))
	dumpBytesArray((*p)[:]) //[hello]
}

func dumpBytesArray(arr []byte) {
	fmt.Printf("[")
	for _, b := range arr {
		fmt.Printf("%c" , b)
	}
	fmt.Println("]")
}
