package main

import "fmt"

func sayHello(fullName string) string {
	return "Hello " + fullName
}

func main(){
	fmt.Println(sayHello("Bangkit Juang Raharjo"))
}