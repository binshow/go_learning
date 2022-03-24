package main

import "fmt"

func main() {
	s := "界"
	fmt.Println(stringToBin(s))
	s = "bin斌"
	fmt.Println(stringToBin(s))

}
// 字符串转换成二进制
func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b",binString, c)
	}
	return
}
