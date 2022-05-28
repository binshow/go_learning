package main

import (
	"flag"
	"fmt"
	"go_learning/instrument_trace/instrumenter"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 实现 Trace 自动注入

var wrote bool

func init() {
	flag.BoolVar(&wrote , "w"  , false , "write result to source file instead of st")
}

func usage() {
	fmt.Println("instrument [-w] xxx.go")
	flag.PrintDefaults()
}

func main() {
	fmt.Println(os.Args)
	flag.Usage = usage
	flag.Parse() // 解析命令行参数

	if len(os.Args) < 2{
		usage()
		return
	}

	var file string
	if len(os.Args) == 3 {
		file = os.Args[2]
	}
	if len(os.Args) == 2 {
		file = os.Args[1]
	}
	if filepath.Ext(file) != ".go"{ // 对源文件扩展名进行校验
		usage()
		return
	}

	var ins instrumenter.Instrumenter
	newSrc, err := ins.Instrument(file)
	if err != nil {
		panic(err)
	}

	if newSrc == nil {
		fmt.Printf("no trace added for %s\n" , file)
		return
	}

	if !wrote {
		fmt.Println(string(newSrc)) // 将生成的新代码内容输出到stdout
		return
	}

	if err = ioutil.WriteFile(file , newSrc , 0666); err != nil{
		fmt.Printf("write %s error : %v\n" , file , err)
		return
	}
	fmt.Printf("instrument trace for %s ok \n" , file)

}
