# Go Goroutine Demonstration

This Go program demonstrates the use of goroutines, a key feature of Go's concurrency model. It launches multiple goroutines to execute a function concurrently and showcases the behavior of goroutines in a simple example.

## Overview

In this program, a function `showNumber` is defined, which prints a number along with a timestamp. The `main` function is responsible for launching multiple goroutines to execute this function concurrently.

## Code Explanation

- `showNumber`: This function takes an integer `num`, converts the current Unix timestamp to a string, and prints both values. It also includes a 100-millisecond sleep to simulate some work.

- `main`: In the `main` function, we do the following:
    - Set the maximum number of CPUs to be used by the Go runtime to 4 using `runtime.GOMAXPROCS(4)`. This ensures that our program can utilize multiple CPU cores.
    - Define the number of iterations, which is set to 10 in this example.
    - Launch goroutines to execute `showNumber` for each iteration.
    - Print "Goodbye!" to the console.
    - Use `runtime.Gosched()` to yield the processor, allowing other goroutines to run.
    - Calculate and print the elapsed time.
