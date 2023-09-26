# Go Concurrency Threads Management

This Go program provides a simple demonstration of how to query and set the number of available threads for Go's concurrency using the `runtime` package.

## Overview

In this program, we explore how to:

- Query the number of available threads using `runtime.GOMAXPROCS(0)`.

- Set the number of available threads using `runtime.GOMAXPROCS(n)`.

## Code Explanation

- `listThreads()`: This function queries and returns the number of available threads using `runtime.GOMAXPROCS(0)`.

- `main`: In the `main` function, we do the following:
    - Set the maximum number of CPUs to be used by the Go runtime to 2 using `runtime.GOMAXPROCS(2)`. This limits the program to use only 2 threads.
    - Call `listThreads()` to query and print the number of available threads.

## Output

The program will display the number of threads available to Go, which may vary depending on your system's configuration and the value set using `runtime.GOMAXPROCS`.
