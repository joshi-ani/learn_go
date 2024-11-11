package main

import "fmt"

func main() {
	// Pointers section

	myInt := 10

	fmt.Println("my integer value is:", myInt)

	myFirstPointer := &myInt

	// points to a location that stores the address where myInt's value 10 is stored in memory.
	fmt.Println("my first pointer address is:", myFirstPointer)

	*myFirstPointer = 15

	fmt.Println("myInt is now:", myInt)

	changeValueOfPointer(&myInt)
	fmt.Println("After function call myInt is now:", myInt)
}

func changeValueOfPointer(num *int) {
	*num = 25
}
