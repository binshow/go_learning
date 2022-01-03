package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	//1. 定义一个 Buffer型变量b , 其中 Buffer 实现了 Reader 和 Writer接口
	var b bytes.Buffer
	//2. 往这个Buffer里面写入字符串
	b.Write([]byte("hello!binshow"))
	//3. 拼接字符串到Buffer中
	fmt.Fprint(&b, ",", "https://binshow.github.io/", "study with me")
	//4. 将这个Buffer输出到标准输出流中
	b.WriteTo(os.Stdout)
}
