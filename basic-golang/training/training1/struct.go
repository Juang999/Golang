package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func main(){
	juang := Customer{
		Name: "Bangkit Juang Raharjo",
		Address: "Bogor",
		Age: 23,
	}

	fmt.Println(juang)
}