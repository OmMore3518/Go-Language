package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func readFromFile() (float64,error){
	byteop, err:=os.ReadFile("balance.txt")
	if(err != nil){
		return 1000, errors.New("Balnce file not found")
	}
	stringop:= string(byteop)
	ans, err := strconv.ParseFloat(stringop,64)
	if(err!=nil){
		return 1000, errors.New("Cannot Converted to float")
	}
	return ans, nil
}

func writeIntoFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}


func main(){
	var choice int
	balance,error:= readFromFile()
	if(error!=nil){
		fmt.Print("Error!!")
		fmt.Println(error)
		fmt.Println("------------------------")
		//panic("Can't Continue")
	}
	var amount float64

    num := true
	//loop :
	for num{
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
			writeIntoFile(balance)
		}else{
			fmt.Println("Please Enter Positive Amount")
		}
	case 3:
		fmt.Println("Please Enter Amount for Withdraw:")
		var wamount float64
		fmt.Scan(&wamount)
		if wamount>0 && wamount<=balance {
			balance -= wamount
			fmt.Println("Balance Withdrawn Successfully")
			writeIntoFile(balance)
		}else{
			fmt.Println("Insufficient Balnace..!!")
		}
	case 4:
		fmt.Println("Thank You! Visit Again!")
		//break loop
		//Only Break Exits the nearest loop like these exit switch not for
		num=false
	default:
		fmt.Println("Please Enter the Correct Option")
	}
	}
	fmt.Print("Thank You For Choosing Our Bank")






// 	bankBanance := 1000
// 	fmt.Println("Welcome to Go Bank")
// 	for{
// 	fmt.Println("Please choose what you want to do:")
// 	fmt.Println("1.Show balance")
// 	fmt.Println("2.Deposit amount")
// 	fmt.Println("3.Winthdraw amount")
// 	fmt.Println("4.Exit")
// 	var choice int
// 	fmt.Scan(&choice)
// 	if(choice==1){
// 		fmt.Println("Your account balance is:",bankBanance)
// 	}else if(choice==2){
// 		fmt.Println("Please Enter amount to deposite:")
// 		var depamt int
// 		fmt.Scan(&depamt)
// 		if (depamt<0){
// 			fmt.Println("Please Enter Valid Amount")
// 			continue
// 		}
// 		bankBanance+=depamt
// 		fmt.Println("Amount successfully deposited!")
// 	}else if (choice==3){
// 		fmt.Print("Please enter amount to withdraw:")
// 		var withamt int
// 		fmt.Scan(&withamt)
// 		if(withamt<0){
// 			fmt.Println("Please Enter Valid Amount")
// 			continue
// 		}
// 		if(withamt>bankBanance){
// 			fmt.Println("Insufficient Balance!")
// 			break
// 		}
// 		withamt-=bankBanance
// 	}else if(choice==4){
// 		fmt.Println("Thank You!")
// 		break
// 	}
// 	}
// 	fmt.Print("Visit Again to Go Bank!!")
 }