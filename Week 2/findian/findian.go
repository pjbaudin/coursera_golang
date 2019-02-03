package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	s "strings"
)

func main() {
	// Prompt user
	fmt.Print("Enter a string: ")

	// Read input from user
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// Text processing
	text = strings.Replace(text, " ", "", -1)
	text = s.ToLower(text)

	// Check string
	// fmt.Print("Your string is: ", text, "\n")                      // Print the string
	// fmt.Print("Your string is: ", string(text[len(text)-3]), "\n") // Print the string

	// Run condition on string
	if string(text[0]) == "i" && string(text[len(text)-3]) == "n" && s.Contains(text, "a") {
		fmt.Print("Found")
	} else {
		fmt.Print("Not Found")
	}
}
