package main

import "fmt"

func getGoodBye(name string) (name1 string) {
	name1 = "goodbye " + name
	return 
}

func main(){
	goodbye := getGoodBye
	fmt.Println(goodbye("Bangkit Juang Raharjo"))
}