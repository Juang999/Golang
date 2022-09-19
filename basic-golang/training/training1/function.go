package main

import "fmt"

func callHello(name string, age int) (string, int) {
	return name, age
}

func main(){
	myName, myAge := callHello("Bangkit Juang Raharjo", 23)

	fmt.Println(myName, myAge)
}