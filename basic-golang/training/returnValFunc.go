package main

import "fmt"

func callName(firstName string) string {
	return "Hello " + firstName
}

func main(){
	// firstName := "Bangkit"
	// value := callName("Bangkit Juang Raharjo")

	fmt.Println(callName("Bangkit Juang Raharjo"))
}