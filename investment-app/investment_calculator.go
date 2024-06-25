package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	// var years float64 = 10

	fmt.Print("Enter the investment amount: ")
	// Use pointer to store the user input value in the investmentAmount variable
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
