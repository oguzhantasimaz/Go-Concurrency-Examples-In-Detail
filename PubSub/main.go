package main

import (
	"fmt"
	"time"
)

var comm = make(chan bool)
var done = make(chan bool)

func producer() {
	for i := 0; i < 10; i++ {
		comm <- true
		time.Sleep(300 * time.Millisecond)
	}
	done <- true
}

func consumer() {
	for {
		communication := <-comm
		fmt.Println("Communication from producer received!", communication)
	}
}

func main() {
	go producer()
	go consumer()
	<-done
	fmt.Println("All Done!")
}
