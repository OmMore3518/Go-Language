package main

import "fmt"

func main(){
	var revenue float64
	var expenses float64
	var taxrate float64

	fmt.Print("Enter the Revenue:")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses:")
	fmt.Scan(&expenses)
	fmt.Print("Enter the taxrate:")
	fmt.Scan(&taxrate)

	EBT := revenue - expenses
	taxAmount := (EBT*taxrate)/100
	profit := EBT - taxAmount
	ratio := EBT/profit

	fmt.Println("EBT:", EBT)
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:",ratio)
}