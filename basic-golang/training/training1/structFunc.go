package main

import "fmt"

type Customer struct {
	Name, Address string
	Age			  int
}

func (customer Customer) sayHi() (realName string, realAge int) {
	realName = customer.Name
	realAge = customer.Age
	
	return
}

func main(){
	juang := Customer{
		Name: "Bangkit Juang Raharjo",
		Age: 23,
	}

	theRealName, theRealAge := juang.sayHi()
	fmt.Println(theRealName)
	fmt.Println(theRealAge)
}