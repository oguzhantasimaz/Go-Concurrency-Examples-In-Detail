package main

//PAGE 19

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func showNumber(num int) {
	tstamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println(num, tstamp)
	time.Sleep(time.Millisecond * 100)
}
func main() {
	var start int64
	start = time.Now().UnixNano()

	runtime.GOMAXPROCS(4)

	iterations := 10
	for i := 0; i <= iterations; i++ {
		go showNumber(i)
	}
	fmt.Println("Goodbye!")
	runtime.Gosched()

	elapsed := float64(time.Now().UnixNano()-start) / 1000000000
	fmt.Println("Elapsed time:", elapsed, "seconds")

}
