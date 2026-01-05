package main

import (
	"fmt"
	"math"
)

func main(){
	const inflationRate = 2.5
	var amountInvested float64
	var returnRate float64
	var years float64

	fmt.Println("Enter the amount that you want to invest")
	fmt.Scan(&amountInvested)
	fmt.Println("Enter the returnRate:")
	fmt.Scan(&returnRate)
	fmt.Println("Enter the Number of Years:");
	fmt.Scan(&years)

	futureValue := amountInvested * math.Pow(1+returnRate/100, years)
	futureRealValue := futureValue * math.Pow(1+returnRate/100,years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
} 