package main

import "fmt"

func callName(firstName string, secondName string){
	fmt.Println("hello", firstName, secondName)
}

func main(){
	firstName := "Bangkit"
	secondName := "Juang"

	callName(firstName, secondName)
}