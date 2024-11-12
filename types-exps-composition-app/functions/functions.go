package main

import "fmt"

// Assign function to type
// First create a type

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// Below function is called a receiver function. In this example it will take the latest values of Animal as a pointer to the struct is passed as argument.
func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println()
}

func (a *Animal) HowManyLegs() {
	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
	fmt.Println()
}

func main() {
	z := addTwoNumber(2, 4)

	fmt.Println("Addition of two numbers is", z)

	myTotal := sumMany(2, 3, 4, 5, 6, -88, 7, 32, -5)

	fmt.Println("My total is:", myTotal)

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumberOfLegs = 4
	dog.Says()
	dog.HowManyLegs()

	cat := Animal{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}
	cat.Says()
	cat.HowManyLegs()
}

// If there is requirement to name the return type then add the return type as (sum int)
// This is called naked return. NOTE: Only use naked return for very short functions because it makes the readability worse.
func addTwoNumber(x, y int) (sum int) {
	sum = x + y
	return
}

// This is called a variadic function that takes any number of arguments.
// NOTE: There can only be one variadic argument in a function otherwise if there are two arguments then the compiler won't know which is which.
func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}

	return total
}
