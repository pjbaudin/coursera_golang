package main

import "fmt"

var t int = 65

func foo() *int {
	a := 1
	// return pointer
	return &a
}

func main() {
	// Define variables
	var x int = 15
	var y int
	var z int
	// define variables for use with pointers
	var ip *int
	var b *int
	var cvar int32 = 1
	var d int16 = 2

	// Use outside function to call for a pointer
	b = foo()
	// Print value pointed at
	fmt.Printf("%d", *b)
	fmt.Println("\n")

	// More example on how to use pointers
	// Copy value in a variable to another using only one pointer variable
	ip = &x
	fmt.Println("X is ", x)
	fmt.Println("My IP is ", ip)
	fmt.Println("y is ", y, "\n")
	fmt.Println("z is ", z, "\n")

	y = *ip
	fmt.Println("X is ", x)
	fmt.Println("My IP is ", ip)
	fmt.Println("Y is ", y, "\n")

	ip = &y
	z = *ip
	fmt.Println("Y is ", y)
	fmt.Println("Z is ", z, "\n")

	// Printing variable outside of function scope
	fmt.Println("t is ", t)

	// Example using print function
	fmt.Printf("Hi \n")
	name := "Joe"
	fmt.Printf("Hi " + name + "\n")

	// Printing format string
	fmt.Printf("Hi %s", name, "\n")

	// assigning and type conversion
	cvar = int32(d)
	fmt.Println("c is: ", cvar)
	fmt.Println("d is: ", d)
}
