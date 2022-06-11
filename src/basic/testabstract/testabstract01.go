package testabstract

import "fmt"

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// Deposit methods
// save
func (account *Account) Deposit(money float64, pwd string) {
	if pwd != account.Pwd {
		fmt.Println("wrong pwd")
	}

	if money <= 0 {
		fmt.Println("wrong money")
	}

	account.Balance += money
}

func (account *Account) GetMoney(money float64, pwd string) {

	if pwd != account.Pwd {
		fmt.Println("wrong pwd")
	}

	if account.Balance <= 0 || money > account.Balance {
		fmt.Println("wrong money")
	}

	account.Balance -= money
}

func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("wrong pwd")
	}

	fmt.Println(account.Balance)
}
