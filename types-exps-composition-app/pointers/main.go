package main

import "fmt"

func main() {
	// Pointers section

	myInt := 10

	fmt.Println("my integer value is:", myInt)

	myFirstPointer := &myInt

	fmt.Println("my first pointer address is:", myFirstPointer)
}
