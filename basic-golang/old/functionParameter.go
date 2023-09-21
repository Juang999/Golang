package main

import "fmt"

type Filter func(string)string

func sayHello(name string, filter Filter) string {
	return "Hello " + filter(name)
}

func thisFilter(name string) string {
	if name == "Bangkit Juang Raharjo" {
		return "mas ganteng" 
	} else if name == "Ismail Nur Alam" {
		return "mas jawa"
	} else {
		return "ihhhh.... wibu"
	}
}

func main(){
	fmt.Println(sayHello("Ismail Nur Alam", thisFilter))
}