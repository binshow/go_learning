package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var(
	n = flag.Bool("n" , false , "omit trailing newline")
	s = flag.String("s" , "" , "separator")
)

var out io.Writer = os.Stdout

func main() {
	flag.Parse() // 解析参数
	if err := echo(!*n , *s , flag.Args()) ; err != nil{
		fmt.Fprintf(os.Stderr , "echo : %v\n" , err)
		os.Exit(1)
	}
}

func echo(newline bool, sep string, args []string) error {
	fmt.Fprint(out , strings.Join(args , sep)) // 向终端写入命令
	if newline {
		fmt.Fprintln(out)
	}
	return nil
}
