package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount, years := 1000.0, 10.0
	expectedReturnRate := 5.5
	// var years float64 = 10

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}
