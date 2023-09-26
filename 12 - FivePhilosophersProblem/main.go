package main

import (
	"fmt"
	"sync"
	"time"
)

const numPhilosophers = 5

var forks [numPhilosophers]*sync.Mutex
var wg sync.WaitGroup

func init() {
	for i := 0; i < numPhilosophers; i++ {
		forks[i] = &sync.Mutex{}
	}
}

func philosopher(id int, leftFork, rightFork *sync.Mutex, eatCount int, done chan bool) {
	defer wg.Done()

	for i := 0; i < eatCount; i++ {
		fmt.Printf("Philosopher %d is thinking\n", id)

		leftFork.Lock()
		rightFork.Lock()

		fmt.Printf("Philosopher %d is eating\n", id)
		time.Sleep(time.Second * 3) // Simulate eating
		fmt.Printf("Philosopher %d is done eating\n", id)

		leftFork.Unlock()
		rightFork.Unlock()
	}

	done <- true
}

func main() {
	done := make(chan bool, numPhilosophers)

	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go philosopher(i+1, forks[i], forks[(i+1)%numPhilosophers], 3, done)
	}

	wg.Wait()

	close(done)
}
