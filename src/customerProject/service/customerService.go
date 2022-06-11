package service

import "firstProject/src/customerProject/model"

type CustomerService struct {
	customers []model.Customer
	// 声明一个字段表示当前切片有多少客户
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "zhangsan", "male", 20, "112", "zs@gmail.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

// Add 这里一定要用引用传递，保证在同一个customer service下面操作
func (cs *CustomerService) Add(customer model.Customer) bool {
	cs.customerNum++
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

func (cs *CustomerService) FindById(id int) int {
	index := -1

	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			index = i
		}
	}
	return index
}

// Delete 这里如果不给传引用类型的话，CustomerService里面的 customer切片不会造成任何影响！
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	if index == -1 {
		return false
	}

	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}
