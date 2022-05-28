package main

import (
	"fmt"
	"reflect"
)

type Interface interface {
	M1()
	M2()
}

type T struct {
}

func (t T) M1() {
}

func (t *T) M2() {
}

func main() {
	var n int
	dumpMethodSet(n) // int's method is empty
	dumpMethodSet(&n) //*int's method is empty


	var t T
	dumpMethodSet(t) //
	dumpMethodSet(&t)

	//main.T's method set :
	//- M1
	//
	//*main.T's method set :
	//- M1
	//- M2

}

// 查看类型的方法集合
func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)
	if dynTyp == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}
	n := dynTyp.NumMethod() // 查看方法集合
	if n == 0{
		fmt.Printf("%s's method is empty\n" , dynTyp)
		return
	}
	fmt.Printf("%s's method set :\n" , dynTyp)
	for i := 0; i < n; i++ {
		fmt.Println("-" , dynTyp.Method(i).Name)
	}
	fmt.Printf("\n")

}