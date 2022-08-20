package main

import "fmt"

func callName() (string, string) {
	return "Bangkit", "Juang"
}

func main(){
	firstName, lastName := callName()

	fmt.Println(firstName)
	fmt.Println(lastName)
}