package main

import "fmt"

type HasName interface {
	GetName() string
}

type Customer struct {
	Name string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (customer Customer) GetName() string {
	return customer.Name
}

func main() {
	customer := Customer{Name: "Bangkit Juang Raharjo"}

	sayHello(customer)
}