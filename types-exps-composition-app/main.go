package main

import "fmt"

// aggregate types (array, struct)

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketOfSeats bool
	Make          string
	Model         string
	Year          int
}

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main() {
	var myStrings [3]string

	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("First element in array is:", myStrings[0])

	// One way to initialize the values of Car struct.
	// var myCar Car
	// myCar.NumberOfTires = 4
	// myCar.BucketOfSeats = true
	// myCar.Luxury = true
	// myCar.Make = "Skoda"
	// myCar.Model = "Slavia"
	// myCar.Year = 2024

	// Another way to initialize the values of Car struct.
	myCar := Car{
		NumberOfTires: 4,
		BucketOfSeats: true,
		Luxury:        false,
		Make:          "Skoda",
		Model:         "Slavia",
		Year:          2024,
	}

	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)

}
