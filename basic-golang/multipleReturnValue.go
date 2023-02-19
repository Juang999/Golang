package main

import "fmt"

func sayHello() (string, string) {
	return "Juang", "Raharjo"
}

func main(){
	namaPertama, namaTerakhir := sayHello()

	fmt.Println(namaPertama)
	fmt.Println(namaTerakhir)
}