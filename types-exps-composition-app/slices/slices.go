package main

import (
	"fmt"
	"sort"
)

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

	// Print length of the slice
	fmt.Println("The slice is: ", len(animals), " elements long")

	// Check if slice is sorted
	fmt.Println("Is is sorted?", sort.StringsAreSorted(animals))

	// If not sorted then sort the slice
	sort.Strings(animals)
	fmt.Println("Is is sorted now?", sort.StringsAreSorted(animals))
	fmt.Println(animals)

	animals = DeleteFromSlice(animals, 1)
	fmt.Println(animals)
}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a
}
