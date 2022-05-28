package main

import (
	"fmt"
	"io"
	"strings"
)

/**
	测试 string.Reader 结构体及其所有的方法
 */

func main() {

	// 基于一个字符串生成一个实现了很多IO接口的阅读器
	reader := strings.NewReader("binshowandustc")


	fmt.Printf("len : %d\n",  reader.Len())  //还未被读的string的长度 ===> 10
	fmt.Printf("size : %d\n", reader.Size()) // 原始string长度，不会被改变的，每次调用都是一样的 ===> 10
	// 自定义实现的String方法
	reader.String() // 自定义Reader的String方法: 字符串 =  binshow , 已读内容长度 =  0

	// Read会影响未读长度的数值，为r.i+n，n为读的长度
	fmt.Println("\nRead 读了5个byte之后的size和len")
	buf1 := make([]byte, 5)
	_, _ = reader.Read(buf1)
	fmt.Println("buffer read : ", string(buf1)) // ===> abcde

	fmt.Printf("len : %d\n",  reader.Len())  //还未被读的string的长度 ===> 5
	fmt.Printf("size : %d\n", reader.Size()) // 原始string长度，不会被改变的，每次调用都是一样的 ===> 10

	// ReadAt不会影响未读长度的数值，为r.i+n，n为读的长度
	s := strings.NewReader("abcdefghij")
	fmt.Println("\nReatAt 读了指定长度之后的字符串，不影响任何len或size")
	buf2 := make([]byte, 6)
	_, _ = s.ReadAt(buf2, 6)

	fmt.Println("buffer read : ", string(buf2)) // ===> ghij

	fmt.Printf("len : %d\n", s.Len()) //还未被读的string的长度 ===> 10

	fmt.Printf("size : %d\n", s.Size()) // 原始string长度，不会被改变的，每次调用都是一样的 ===> 10


	// ReadByte每次只读1byte，从未读的index开始，每次只返回1byte，相对应的读完之后r.i会+1
	s = strings.NewReader("abcdefghij")
	fmt.Println("\nReadByte 只读1个byte，len修改 size不修改")

	buf3, _ := s.ReadByte()

	fmt.Println("buffer read : ", string(buf3)) // ===> a

	fmt.Printf("len : %d\n", s.Len()) //还未被读的string的长度 ===> 9

	fmt.Printf("size : %d\n", s.Size()) // 原始string长度，不会被改变的，每次调用都是一样的 ===> 10

	// UnreadByte，从r.i开始，往后退，r.i会-1
	fmt.Println("\nUnreadByte len多1，不影响任何len或size")

	_ = s.UnreadByte()

	fmt.Printf("len : %d\n", s.Len()) //还未被读的string的长度 ===> 10

	fmt.Printf("size : %d\n", s.Size()) // 原始string长度，不会被改变的，每次调用都是一样的 ===> 10

	// Seek，算偏移，可以指定到那个index，会改变len
	fmt.Println("\nSeek 偏移位数，一般配合read来用")

	bias, _ := s.Seek(4, io.SeekCurrent)

	buf4 := make([]byte, 3)
	_, _ = s.ReadAt(buf4, bias)

	fmt.Println("buffer read : ", string(buf4)) // ===> efg

	fmt.Printf("len : %d\n", s.Len()) //还未被读的string的长度 ===> 6

	fmt.Printf("size : %d\n", s.Size()) // 原始string长度，不会被改变的，每次调用都是一样的 ===> 10

	bias, _ = s.Seek(-2, io.SeekCurrent)

	buf5 := make([]byte, 3)
	_, _ = s.ReadAt(buf5, bias)

	fmt.Println("buffer read : ", string(buf5)) // ===> cde

	fmt.Printf("len : %d\n", s.Len()) //还未被读的string的长度 ===> 8

	fmt.Printf("size : %d\n", s.Size()) // 原始string长度，不会被改变的，每次调用都是一样的 ===> 10

	zw := strings.NewReader("测a中文")

	ch, size, _ := zw.ReadRune()

	fmt.Printf("buffer read : %c \n", ch) // ===> 测

	fmt.Println("buffer read  size : ", size) // ===> 3

	ch, size, _ = zw.ReadRune()

	fmt.Printf("buffer read : %c \n", ch) // ===> a

	fmt.Println("buffer read  size : ", size) // ===> 1

	_ = zw.UnreadRune()

	ch, size, _ = zw.ReadRune()

	fmt.Printf("buffer read : %c \n", ch) // ===> a

	fmt.Println("buffer read  size : ", size) // ===> 1



}
