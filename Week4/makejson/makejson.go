package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Write a program which prompts the user
// to first enter a name, and then enter an address.
// Your program should create a map and add the name
// and address to the map using the keys “name” and “address”,
// respectively.
// Your program should use Marshal() to create a JSON object
// from the map, and then your program should print the JSON object.

func main() {
	// Prompt user to enter name
	fmt.Print("Enter a name: ")

	// Read input from user
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	// Prompt user to enter name
	fmt.Print("Enter an address: ")
	// Read user input
	address, _ := reader.ReadString('\n')

	// Create map
	m := make(map[string]string)

	// Assign user input
	m["Name"] = name
	m["Address"] = address

	fmt.Println("What you entered:", m)

	b, _ := json.Marshal(m)
	fmt.Println("JSON object created with Marshall function: ", b)

}
