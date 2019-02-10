package main

import "fmt"
import "sort"
import "strconv"

func main() {
	// Define user input variables
	var input string

	// Define slice
	userSlice := make([]int, 1)

	for input != "x" {
		// Prompt user to enter integer
		fmt.Print("Press x to exit of Enter an integer: ")

		// Read user input
		fmt.Scanln(&input)

		// Convert input to intuser
		intuser, err := strconv.Atoi(input)
		if err != nil {
			// handle error
		}

		// Loop to compare value
		for i, v := range userSlice {
			if v == intuser {
				userSlice[i] = v
				break
			} else {
				userSlice = append(userSlice, intuser)
				break
			}
		}

		sort.Ints(userSlice)
		// Print user input
		fmt.Print(userSlice)
		fmt.Printf("len=%d", len(userSlice))
	}
}
