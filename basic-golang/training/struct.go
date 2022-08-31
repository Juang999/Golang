package main

import "fmt"

type Contact struct {
	firstName, lastName, phoneNumber string
}

func main() {
	contactPerson := Contact{"Bangkit", "Juang", "081325507780"}
	secondContactPerson := Contact{"Mujahid", "Muslim", "08123456789"}

	fmt.Println(contactPerson)
	fmt.Println(secondContactPerson)
}