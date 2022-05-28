package main

import (
	"fmt"
	"os"
	"sync"
)

func main()  {

	var wg sync.WaitGroup
	fmt.Println(wg)

	os.Stat("")

}
