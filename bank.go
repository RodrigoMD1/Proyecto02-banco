package main

import "fmt"

func main() {



	fmt.Println("welcome to go bank")
	fmt.Println(("what do you want to do"))
	fmt.Println("1. check balance")
	fmt.Println("2. deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")



	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)


	fmt.Println("Your choice is", choice)

}
