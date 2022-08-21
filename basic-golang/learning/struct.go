package main

import "fmt"

type Customer struct {
	Name, Address string
	Age			  int
}

func main() {
	var juang Customer
	juang.Name = "Bangkit Juang Raharjo"
	juang.Address = "Kuvuki land"
	juang.Age = 23

	fmt.Println(juang)

	latami := Customer{
		Name: "Rahmat Bagus Latami",
		Address: "Indonesia",
		Age: 23,
	}

	fmt.Println(latami)
}