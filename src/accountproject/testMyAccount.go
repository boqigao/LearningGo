package main

import (
	"fmt"
)

func main() {
	key := ""
	loop := true
	balance := 10000.0
	money := 0.0
	note := ""
	details := "inAndOut\tcurrentMoney\tinAndOutMoney\tnote\n"

	for {
		fmt.Println("----------Account Software----------")
		fmt.Println("1. Details")
		fmt.Println("2. memo income")
		fmt.Println("3. memo cost")
		fmt.Println("4. exit")
		fmt.Println("please enter 1 ~ 4")

		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("----------Details until now----------")
			fmt.Println(details)
		case "2":
			fmt.Println("memo income")
			fmt.Println("input amount of income:")
			fmt.Scanln(&money)
			fmt.Println("note of income:")
			fmt.Scanln(&note)
			balance += money
			details += "income\t" + fmt.Sprintf("%f", balance) + "\t" + fmt.Sprintf("%f", money) + "\t" + note + "\n"
		case "3":
			fmt.Println("memo cost")
			fmt.Println("memo cost")
			fmt.Println("input amount of cost:")
			fmt.Scanln(&money)
			fmt.Println("note of cost:")
			fmt.Scanln(&note)
			balance -= money
			details += "cost\t" + fmt.Sprintf("%f", balance) + "\t" + fmt.Sprintf("%f", money) + "\t" + note + "\n"
		case "4":
			loop = false
		default:
			fmt.Println("please enter 1 ~ 4")
		}
		if !loop {
			break
		}
	}

	fmt.Println("You have logged out ")
}
