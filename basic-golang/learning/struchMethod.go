package main

import "fmt"

type Customer struct {
	Name, Address string
	Age 		  int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("hello", customer.Name, "my name is", name)
}

func main() {
	var juang Customer
	juang.Name = "Bangkit Juang Raharjo"
	juang.Address = "Bogor"
	juang.Age = 23

	juang.sayHi("Mujahid")
}