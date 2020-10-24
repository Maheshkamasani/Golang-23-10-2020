package main 

import "fmt"

// main function 
func main() { 

	// Array of odd numbers 
	odd := [7]int{1, 3, 5, 7, 9, 11, 13} 

	// using range keyword with for loop to 
	// iterate over the array elements 
	for i, item := range odd { 

		// Prints index and the elements 
		fmt.Printf("odd[%d] = %d \n", i, item) 
	} 
} 
