package main

import (
	"firstProject/src/customerProject/model"
	"firstProject/src/customerProject/service"
	"fmt"
)

type customerView struct {
	key  string
	loop bool

	customerService *service.CustomerService
}

// show all the customers

func (cv customerView) list() {
	customers := cv.customerService.List()
	fmt.Println("----------Customers List----------")
	fmt.Println("id\tname\tgender\tage\tphone\tmail")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-------Customers List Down--------")
}

// show menu

func (cv customerView) mainMenu() {
	for {
		fmt.Println("----------Customer Manager----------")
		fmt.Println("1. add customer")
		fmt.Println("2. update customer")
		fmt.Println("3. delete customer")
		fmt.Println("4. list customer")
		fmt.Println("5. exit")
		fmt.Println("input 1~5")

		fmt.Scanln(&(cv.key))
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			fmt.Println("update customer")
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.loop = false
		}

		if !cv.loop {
			break
		}
	}

	fmt.Println("you quit the customer management system")
}

func (cv customerView) add() {
	fmt.Println("----------Add customer----------")

	fmt.Println("name: ")
	name := ""
	fmt.Scanln(&name)

	fmt.Println("gender:")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Println("age:")
	age := 0
	fmt.Scanln(&age)

	fmt.Println("phone:")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Println("mail:")
	mail := ""
	fmt.Scanln(&mail)

	// create a new customer
	customer := model.NewCustomer2(name, gender, age, phone, mail)
	if cv.customerService.Add(customer) {
		fmt.Println("---------------- add successfully -------")
	} else {
		fmt.Println("------------- add failed- ----------------")
	}
}

func (cv customerView) delete() {
	fmt.Println("------------ delete customers ----------")
	fmt.Println("please input id (input -1 to quit):")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	fmt.Println("really delete?")
	choice := ""
	fmt.Scanln(&choice)

	if choice == "Y" || choice == "y" {
		if cv.customerService.Delete(id) {
			fmt.Println("Successfully deleted!")
		} else {
			fmt.Println("id not existing!")
		}
	}

}

func main() {
	var customerView = customerView{
		key:  "",
		loop: true,
	}

	customerView.customerService = service.NewCustomerService()

	customerView.mainMenu()
}
