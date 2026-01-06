package main

import "fmt"

func main(){
	var choice int
	balance:= 1000
	var amount int

	for {
    fmt.Println("Welcome to Go Bank!")
	fmt.Println("Please enter your choice:")
	fmt.Println("1.Check Balance")
	fmt.Println("2.Add Amount")
	fmt.Println("3.Withdraw Amount")
	fmt.Println("4.Exit")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Balance:",balance)
	case 2:
		fmt.Print("Enter the amount to add:")
		fmt.Scan(&amount)
		if amount>0{
			balance+=amount
			fmt.Println("Amount Successfully Added")
		}else{
			fmt.Println("Please Enter Positive Amount")
		}
	case 3:
		fmt.Println("Please Enter Amount for Withdraw:")
		var wamount int
		fmt.Scan(&wamount)
		if wamount>0 && wamount<=balance {
			balance -= wamount
			fmt.Println("Balance Withdrawn Successfully")
		}else{
			fmt.Println("Insufficient Balnace..!!")
		}
	case 4:
		fmt.Println("Thank You! Visit Again!")
		return
	default:
		fmt.Println("Please Enter the Correct Option")
	}
	}
}