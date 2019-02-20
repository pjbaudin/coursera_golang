package main

import "fmt"

// Write a program which prompts the user
// to first enter a name, and then enter an address.
// Your program should create a map and add the name
// and address to the map using the keys “name” and “address”,
// respectively.
// Your program should use Marshal() to create a JSON object
// from the map, and then your program should print the JSON object.

// Log structure to collect name and address of user input
type Log struct {
	name, address string
}

var m map[string]Log

func main() {
	// Define user input variables
	var name string
	var address string

	// Prompt user to enter name
	fmt.Print("Enter a name: ")
	// Read user input
	fmt.Scanln(&name)

	// Prompt user to enter name
	fmt.Print("Enter an address: ")
	// Read user input
	fmt.Scanln(&address)

	m = make(map[string]Log)
	m[name] = Log{name, address}

	fmt.Println(m[name])
}
