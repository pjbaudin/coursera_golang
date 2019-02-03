package main

import "fmt"
import "sort"

func main() {
	// Define user input variables
	var input int
	// Define slice
	userSlice := make([]int, 3)

	for i := 0; i < 4; i++ {
		// Prompt user to enter integer
		fmt.Print("Press x to exit of Enter an integer: ")

		// Read user input
		fmt.Scanln(&input)

		// Loop to compare value
		for i, v := range userSlice {
			if v <= input {
				userSlice[i] = input
				break
			} else {
				userSlice = append(userSlice, input)
				break
			}
		}

		sort.Ints(userSlice)

	}

	// Print user input
	fmt.Print(userSlice)
	fmt.Printf("len=%d", len(userSlice))
}
