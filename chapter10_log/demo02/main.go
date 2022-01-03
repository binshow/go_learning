package main

import (
	"io"
	"log"
	"os"
)

// 自定义不同级别的日志
var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	Info = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	// MultiWriter 表示可以在多个输出流中写入
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {

	Info.Println("binshow:", "https://binshow.github.io/")
	Warning.Printf("binshow:：%s\n", "https://binshow.github.io/")
	Error.Println("测试一下error")

}
