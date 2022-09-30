package main

import "fmt"

func testing1(param interface{}) interface{} {
	return param
}

func main() {
	fmt.Println(testing1("Bangkit Juang Raharjo"))
}