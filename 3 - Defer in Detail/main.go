package main

import "fmt"

func main() {
	aValue := new(int)

	//defer fmt.Println(*aValue)
	// If here returns 0

	for i := 0; i < 100; i++ {
		*aValue++
		defer fmt.Println(*aValue)
		// If here returns 100,99,98,97,96,95,94,93,92,91,90,89,88,87,86,85,84,83,82,81...1,0
	}

	// defer fmt.Println(*aValue)
	// If here returns 100

}
