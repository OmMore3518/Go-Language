package main

import "fmt"

func main(){
	var revenue float64
	var expenses float64
	var taxrate float64

	//fmt.Print("Enter the Revenue:")
	outPut("Enter the Revenue:")
	fmt.Scan(&revenue)
	//fmt.Print("Enter the expenses:")
	outPut("Enter the Expenses:")
	fmt.Scan(&expenses)
	//fmt.Print("Enter the taxrate:") 
	outPut("Enter the taxrate:")
	fmt.Scan(&taxrate)

	// EBT := revenue - expenses
	// taxAmount := (EBT*taxrate)/100
	// profit := EBT - taxAmount
	// ratio := EBT/profit
	EBT,profit,ratio:=calculations(revenue,expenses,taxrate)


	//fmt.Println("EBT:", EBT)
	//fmt.Println("Profit:", profit)
	//fmt.Println("Ratio:",ratio)
	//fmt.Printf("Ebt: %v \ntaxAmount: %v \nprofit %v \nratio %v\n",EBT,taxAmount,profit,ratio)


	//fmt.Printf("EBT:%.2f\ntaxAmount: %.1f\nprofit:%.0f\nratio:%.3f",EBT,taxAmount,profit,ratio)
	formattedOP:=fmt.Sprintf("EBT: %f\n Profit: %f\n Ratio: %f\n ", EBT,profit,ratio)
	fmt.Print(formattedOP)
}
func outPut(text string){
	fmt.Print(text)
}

func calculations(revenue float64, expenses float64, taxrate float64) (float64,float64,float64){
	ebt:=revenue - expenses
	ta:= (ebt*taxrate)/100
	p:= ebt - ta
	r:= ebt/p
	return ebt,p,r
}