package main

import (
	"fmt"
	"time"
)

type user struct{
	firstName string
	middleName string
	lastName string
	createdAt time.Time
}
func (u user) outputUserData(){
	fmt.Print(u.firstName, u.lastName, u.createdAt)
}
func (u *user) updateUser(){
	u.firstName=""
	u.lastName=""
}

func main(){
	userFirstName:=userData("Enter First Name:")
	userMiddleName:=userData("Enter Middle Name:")
	userLastName:=userData("Enter Last Name:")
	//fmt.Print(firstName, middleName, LastName)

	var appUser user 

    appUser = user{
	firstName: userFirstName,
	middleName: userMiddleName,
	lastName: userLastName,
	createdAt : time.Now(),
}
appUser.outputUserData()
appUser.updateUser()
appUser.outputUserData()
}

func userData(value string) string{
	fmt.Print(value)
	var Name string
	fmt.Scan(&Name)
	return Name
}
