package main

import (
	"fmt"
	"sync"
)

func main() {
	current := 0
	iterations := 100
	wg := new(sync.WaitGroup)
	mutex := new(sync.Mutex)
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		go func() {
			mutex.Lock()
			fmt.Println(current)
			current++
			mutex.Unlock()
			fmt.Println(current)
			wg.Done()
		}()
	}
	wg.Wait()
}
