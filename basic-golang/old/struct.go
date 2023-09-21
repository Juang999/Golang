package main

import "fmt"

type Customer struct {
	Name, Address string
	Age 			int
}

func (customer Customer) sayHello() {
	fmt.Println("Hallo", customer.Name)
}

func main(){
	var customer Customer

	customer.Name = "Bangkit Juang Raharjo"
	customer.Address = "Bogor"
	customer.Age = 23

	fmt.Println(customer)

	customer.sayHello()
}