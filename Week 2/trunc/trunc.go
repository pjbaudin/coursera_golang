package main

import "fmt"

func main() {
	fmt.Print("Enter a float: ") //Print function is used to display output in same line
	var x float32
	fmt.Scanln(&x)                        // Take input from user
	fmt.Print("Your float is: ", x, "\n") // Print the float
	fmt.Printf("Your float truncated is: %.0f", x)
}
