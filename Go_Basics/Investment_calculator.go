package main

import (
	"fmt"
	"math"
)

func main(){
	var amountInvested = 1000
	var returnRate = 5.5
	var years = 10

	var futureValue = float64(amountInvested) * math.Pow(1+returnRate/100, float64(years))
	fmt.Print(futureValue)
} 