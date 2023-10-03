package main

import (
	"fmt"
	"time"
)

func main() {
	// Go routine
	go func() {
		fmt.Println("Simulating work")
	}()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Enough's enough")
	}
}
