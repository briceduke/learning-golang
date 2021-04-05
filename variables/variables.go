package main

import "fmt"

var test, test2 string = "test", "test 2"

func main() {
	// Explicit
	var name string = "Brice"
	fmt.Println(name)

	// Implicit
	var last = "Duke"
	fmt.Println(last)

	// Errors
	// var middle string
	// fmt.Println(middle)

	// Use variable that was defined outside of a function
	fmt.Println(test)
	fmt.Println(test2)

	// Shorthand must be in a function
	short := "shorthand"
	fmt.Println(short)
}