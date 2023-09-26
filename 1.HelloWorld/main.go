package main

import (
	"fmt"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
}
func main() {
	// Define hello and world jobs
	hello := new(Job)
	world := new(Job)

	// Set min and max and text for hello
	hello.text = "hello"
	hello.i = 0
	hello.max = 3

	// Set min and max and text for world
	world.text = "world"
	world.i = 0
	world.max = 5

	// Run outputText concurrently for hello
	// For World don't make it concurrent
	go outputText(hello)
	outputText(world)
}
