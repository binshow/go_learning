package main

import (
	"sync"
)

var (
	lockEsKey   = "dups_update_es"
	lockEsValue = "dups_update_es_机器id_当前时间"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
	}
	wg.Wait()
}
