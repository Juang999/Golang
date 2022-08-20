package main

import "fmt"

func main(){
	type NoKTP string

	var noktp NoKTP = "123456789"
	fmt.Println(noktp)
}