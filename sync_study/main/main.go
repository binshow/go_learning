package main

import (
	"fmt"
	"sync"
)

/**
	研究 sync 包中的同步原语
 */

func main() {
				 // 01  |  11   =  11 说明有一个是1就为1
 	fmt.Println(1 | 3)  //3
	fmt.Println(1 | 1)  //1
	fmt.Println(1 | 0)  //1
	fmt.Println(0 | 0)  //0
	fmt.Println(0 | 1)  //1

}



func MutexDemo() {
	wg := sync.Mutex{}
	wg.Lock()

	m := wg

	fmt.Println(m)
}
