package main

import "fmt"

type SayHi interface {
	GetNickName() string
}

type RequestCustomer struct {
	NickName string
}

func sayHello3(sayHi SayHi){
	fmt.Println("hello", sayHi.GetNickName())
}

func (request RequestCustomer) GetNickName() string {
	return request.NickName
}

func main() {
	var juangRaharjo RequestCustomer
	juangRaharjo.NickName = "Juang"

	sayHello3(juangRaharjo)
}