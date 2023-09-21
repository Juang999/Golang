package main

import "fmt"

func sayHello(firstName string, lastName string) (namaPertama, namaTerakhir string) {
	namaPertama = firstName
	namaTerakhir = lastName

	return firstName, lastName
}

func main(){
	thisFirstName, thisLastName := sayHello("Bangkit Juang", "Raharjo")

	fmt.Println(thisFirstName)
	fmt.Println(thisLastName)
}