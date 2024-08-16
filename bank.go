package main

import "fmt"


func main(){
	var accountBalance float64= 10000
	
	fmt.Println("Welcome to Go bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. withdraw money")
	fmt.Println("4. Exit ")
				
	var choice int
	fmt.Println("Your Choice: ")
	fmt.Scan(&choice)
	
	//wantsCheckBalance :=choice == 1
	
	if  choice == 1{
		fmt.Println("your balance is :",accountBalance)
	}else if choice == 2 {
		fmt.Println("Your deposit :")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance = accountBalance + depositAmount
		fmt.Println("Balance updated! new Amount after deposit is :",accountBalance)	
	}else if choice == 3{
		fmt.Println("Withdraw money:")
		var withdrawmoney float64
		fmt.Scan(&withdrawmoney)
		accountBalance -=withdrawmoney
		fmt.Println("Balance updated! new amount after withdrawing the money is :",accountBalance)	
	} else{
		fmt.Println("Goodbye")
	}
	fmt.Println("Your choice:",choice)
}