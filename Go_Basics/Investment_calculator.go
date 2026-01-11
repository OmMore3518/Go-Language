package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main(){

	// var amountInvested float64
	// var returnRate float64
	// var years float64
	// fmt.Println("Enter the amount that you want to invest")
	// fmt.Scan(&amountInvested)
	// fmt.Println("Enter the returnRate:")
	// fmt.Scan(&returnRate)
	// fmt.Println("Enter the Number of Years:")
	// fmt.Scan(&years)

	amountInvested,err1 := userInput("Enter the amount you want to invest")
	// if(err != nil){
	// 	fmt.Print(err)
	// 	return
	// }
	returnRate,err2 := userInput("Enter the Return Rate:")
	// if(err != nil){
	// 	fmt.Print(err)
	// 	return
	// }
	years, err3 := userInput("Enter the number of Years:")
	if(err1 != nil || err2 != nil || err3 != nil){
		fmt.Print(err1)
		return
	}

	futureValue,futureRealValue:= calculations(amountInvested,returnRate,years)
	// futureValue := amountInvested * math.Pow(1+returnRate/100, years)
	// futureRealValue := futureValue * math.Pow(1+returnRate/100,years)

	fmt.Println(futureValue)
	fmt.Printf("%.0f",futureRealValue)
	writeToFile(amountInvested,returnRate,years)
} 


func writeToFile(amountInvested, returnRate, years float64){
	cal:=fmt.Sprintf("Amount: %.1f\n returnRate:%.1f\n years: %.1f\n",amountInvested,returnRate,years)
	os.WriteFile("calulation", []byte(cal), 0644)
}

func userInput(text string) (float64,error){
	var ans float64
	fmt.Println(text)
	fmt.Scan(&ans)
	if(ans==0 || ans<0){
		return 0,errors.New("Value must be positive number")
	}
	return ans, nil
}
func calculations(amountInvested float64, returnRate float64, years float64)(fv float64, rfv float64){
	fv= amountInvested * math.Pow(1+returnRate/100, years)
	rfv= fv * math.Pow(1+returnRate/100,years) 
	return fv,rfv
}