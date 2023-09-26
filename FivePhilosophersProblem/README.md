# Dining Philosophers Problem in Go

The Dining Philosophers problem is a classic synchronization and concurrency problem that illustrates issues in resource allocation and deadlock avoidance. This Go program demonstrates the solution to the Dining Philosophers problem using goroutines and mutex locks.

## Problem Description

The problem involves a scenario where five philosophers sit around a circular dining table, and there are five forks between them. Each philosopher thinks and eats. To eat, a philosopher needs two forks, one on the left and one on the right. Philosophers follow these rules:

1. A philosopher can only eat when both forks are available.

2. A philosopher must put down both forks after eating.

3. Philosophers must think for a while before trying to eat again.

The goal is to ensure that philosophers can eat without causing deadlock, where all philosophers are waiting for a fork held by another philosopher.

## Code Explanation

- `numPhilosophers` represents the number of philosophers, which is set to 5.

- An array of mutex locks, `forks`, is created to represent the forks on the table.

- Each philosopher is implemented as a goroutine, and they follow the rules of picking up forks, eating, and putting down forks.

- A channel `done` is used to signal when each philosopher has finished eating.

## Output

You will see the philosophers engaging in thinking and eating cycles while respecting the rules of the Dining Philosophers problem. The output will demonstrate how mutex locks are used to synchronize access to the forks and ensure that philosophers can eat without causing deadlock.

## Conclusion

This program illustrates a solution to the Dining Philosophers problem using Go's concurrency features. It serves as an example of how mutex locks can be used to manage access to shared resources in a concurrent system.

Feel free to explore and modify this code to experiment with different scenarios or add enhancements.
