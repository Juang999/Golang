package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Bogor", "West Java", "Indonesia"}
	address2 := &address1

	address2.City = "Jakarta"

	*address2 = Address{"Bekasi", "Jawa Timur", "Indonesia"}
	
	fmt.Println(address1)
	fmt.Println(address2)
} 



