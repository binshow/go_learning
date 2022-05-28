package main

import (
	"fmt"
	"io"
	"os"
)

var (
	myFprintf = func(w io.Writer , format string , a ...interface{})(int , error) {
		return fmt.Fprintf(w , format , a)
	}
)

func main() {
	fmt.Printf("%T\n" , myFprintf)
	myFprintf(os.Stdout , "%s\n" , "hello,go")
}
