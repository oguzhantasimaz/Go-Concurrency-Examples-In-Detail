# Understanding Go Deferred Functions

This short Go code snippet demonstrates the behavior of deferred functions in Go, specifically when used within a loop. The code is designed to help you understand how the `defer` statement works in different contexts.

## Code Explanation

In this code snippet, we have a simple loop that increments the value of a variable `aValue` and defers the printing of `*aValue`. The `defer` statement postpones the execution of the `fmt.Println(*aValue)` until the surrounding function (`main` in this case) exits.

Here's what you can observe in this code:

1. `aValue` is initially set to a new integer pointer (`new(int)`).

2. Inside the loop that runs 100 times, `*aValue` is incremented by 1 for each iteration, and the value is printed using `defer`.

3. The code also includes comments indicating what value of `*aValue` you can expect at specific points in the code.
