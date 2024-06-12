package main

import (
	"fmt"
	"myscopeapp/packageone"
)

// package level variable
var one = "One"

// block level variables
func main() {
	// This is a variable shadowing and don't use this
	// var one = "This is a block level variable"

	var somethingElse = "This is a block level variable"
	fmt.Println(somethingElse)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println(newString)

	packageone.Exported(one, somethingElse)
}

func myFunc() {
	// var one = "The number one"

	fmt.Println(one)
}
