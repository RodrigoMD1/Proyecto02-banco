package main

import (
	"examples.com/bank/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

// constante global

const accountBalanceFile = "balance.txt"

////////////////////////////////////////////////////////////////////////////////

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------------")
		//* panic("cant continue without balance file") lo mismo que el return pero da como una señal mas clara de que hay un error
		// return si quiere que el programa no funcione ya sin ese archivo txt por que es obligatorio tenerlo solo tengo que permitir este return
	}

	fmt.Println("welcome to go bank")
	fmt.Println("reach us 24/7", randomdata.PhoneNumber())

	for {
		//* dejando el bucle for solo hace que el codigo se repita infinitamente hasta que se cumpla la condicion de salida que en este caso es el else
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

		switch choice {

		case 1:

			fmt.Println(("Your balance is"), accountBalance)

			///////////////////////////////////////////////////////////////////////////
		case 2:

			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("invalid amount. must be greater than 0.")
				//return // con este return si da 0 entonces detiene el codigo de abajo
				continue //* me permite que si da 0 entonces no se ejecute el codigo de abajo y se vuelve a el principio con el return de antes solo terminaba el programa si daba menos de 0 pero este me lo vuelve a reiniciar
			}

			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New Amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

			///////////////////////////////////////////////////////////////////////////
		case 3:

			fmt.Print("Your withdraw: ")
			var witdrawAmount float64
			fmt.Scan(&witdrawAmount)

			if witdrawAmount <= 0 {
				fmt.Println("invalid amount. must be greater than 0.")
				//return //* con este return si da 0 entonces detiene el codigo de abajo
				continue
			}
			if witdrawAmount > accountBalance {
				fmt.Println("invalid amount. you can't withdraw more than your balance")
				continue
			}
			accountBalance -= witdrawAmount
			fmt.Println("Balance updated! New Amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

			///////////////////////////////////////////////////////////////////////////

		default:

			fmt.Println(("Goodbye!"))
			fmt.Println("thanks for choosing our bank")

			return
			//break //* sirve para salir del bucle for y terminar el programa deja de funcionar si esta dentro de un switch por que en ese caso se tiene que usar un return

		}
	}

}
