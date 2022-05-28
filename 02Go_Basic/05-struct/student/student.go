package student

import "fmt"

type Student struct {
	Class string
	grade int
}


type teacher struct {
	major string
	grade int
}

func Test() {
	s := Student{Class: "aa"}
	fmt.Println(s)
}
