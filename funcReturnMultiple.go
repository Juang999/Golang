package main

import "fmt"

func getFullName() (string, string) {
	return "Bangkit", "Juang"
}

func main(){
	firstName, lastName := getFullName()

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(getFullName())
}