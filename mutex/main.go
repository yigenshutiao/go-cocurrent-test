package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}

	wg.Wait()
	// 不加锁导致的后果, 理论上要:1000000, 实际上给了:650497
	fmt.Println(count)
}
