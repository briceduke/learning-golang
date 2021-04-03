package main

import "fmt"

func main()  {
	// Prints a line, adds \n
	fmt.Println("Hello Line!")
	// Prints a line without \n
	fmt.Print("Hello World!\n")
	// Prints with args, %s means arg type must be a string, %d means int
	fmt.Printf("Hello, my name is %s. It is April %d", "Brice", 3)

	// FPrint prints to an external source
	// SPrint stores output in character buffer
}