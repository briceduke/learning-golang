package main

import "fmt"

func main() {
	city := "Austin"

	switch city {
	case "Austin":
		fmt.Println("Austin")
	case "Dallas", "Fort Worth":
		fmt.Println("DFW")
	case "San Antonio":
		fmt.Println("San Antonio")
	}

	var i int

	switch  {
	case i> 10:
		fmt.Println(" > 10")
	case i != 42:
		fmt.Println("AHHH")
		fallthrough // Stops checking statements below
	default:
		fmt.Println("LT10")
	}
}