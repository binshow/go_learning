package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	court := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)


	go func() {
		play("binshow" , court)
		wg.Done()
	}()

	go func() {
		play("zkd" , court)
		wg.Done()
	}()

	court <- 1
	wg.Wait()
}



func play(name string, court chan int) {
	for {

		ball , ok := <- court
		if !ok {
			fmt.Printf("%s win!!\n" , name)
			return
		}

		if rand.Intn(100) % 13 == 0 {
			fmt.Printf("%s missed!!\n" , name)
			close(court)
			return
		}

		fmt.Printf("%s hit %v!!\n" , name , ball)
		ball++

		court <- ball

	}
}