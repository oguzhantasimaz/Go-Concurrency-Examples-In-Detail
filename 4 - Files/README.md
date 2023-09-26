# Go Concurrent File Output Example

This Go code snippet demonstrates how to use Go's concurrency features to concurrently write text to files. It provides an example of creating two separate goroutines to write "hello" and "world" to their respective text files concurrently.

## Code Explanation

In this code snippet, we have the following components:

- `Job` struct: This struct holds information about a job, including an index `i`, a maximum count `max`, and a text string.

- `outputText` function: This function simulates a job by appending the text to a string and writing it to a file. It does this in a loop, introducing a small delay using `time.Sleep` to mimic work being done.

- In the `main` function, we create two `Job` instances: `hello` and `world`, each with its own text and maximum count.

- We then use the `go` keyword to launch two goroutines that concurrently call the `outputText` function for `hello` and `world`.
