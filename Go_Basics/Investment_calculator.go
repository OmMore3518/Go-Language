package main

import (
	"fmt"
	"math"
)
	
func main(){
	// var amountInvested float64
	// var returnRate float64
	// var years float64
	// fmt.Println("Enter the amount that you want to invest")
	// fmt.Scan(&amountInvested)
	// fmt.Println("Enter the returnRate:")
	// fmt.Scan(&returnRate)
	// fmt.Println("Enter the Number of Years:");
	// fmt.Scan(&years)

	amountInvested := userInput("Enter the amount you want to invest")
	returnRate := userInput("Enter the Return Rate:")
	years := userInput("Enter the number of Years:")

	futureValue,futureRealValue:= calculations(amountInvested,returnRate,years)
	// futureValue := amountInvested * math.Pow(1+returnRate/100, years)
	// futureRealValue := futureValue * math.Pow(1+returnRate/100,years)

	fmt.Println(futureValue)
	fmt.Printf("%.0f",futureRealValue)
} 
func userInput(text string) float64{
	var ans float64
	fmt.Println(text)
	fmt.Scan(&ans)
	return ans
}
func calculations(amountInvested float64, returnRate float64, years float64)(fv float64, rfv float64){
	fv= amountInvested * math.Pow(1+returnRate/100, years)
	rfv= fv * math.Pow(1+returnRate/100,years)
	return fv,rfv
}