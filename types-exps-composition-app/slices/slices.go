package main

import "fmt"

// Slices

func main() {
	var animals []string

	// One way to add elements to the slice
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")
	animals = append(animals, "cow")

	fmt.Println("Animals values:", animals)

	// Another way to print the values one-by-one is to iterate over reference type

	// for _, x := range animals {
	// 	fmt.Println(x)
	// }

	// Notice the change, as we replaced _ with i and that will print the values with a sequential number
	for i, x := range animals {
		fmt.Println(i, x)
	}

	// Print first element
	fmt.Println("First elements: ", animals[0])

	// Print first two elements
	fmt.Println("First two elements are: ", animals[0:2])
}
