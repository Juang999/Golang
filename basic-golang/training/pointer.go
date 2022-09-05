package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeRegionToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Bogor", "West Java", "Indonesia"}
	address2 := &address1

	address2.City = "Jakarta"

	*address2 = Address{"Banjar", "West Java", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	var address = Address{
		City: "Bogor",
		Province: "West Java",
		Country: "",
	}

	changeRegionToIndonesia(&address)

	fmt.Println(address)
}