package main

import "fmt"

type Filter (string) string

func seyHello(name string, filter Filter) string {
	filtered := filter(name)
	return "hello " + filtered
}

func theFilter(name string) string {
	if name == "Babi" {
		return "..."
	} else {
		return name
	}
}

func main(){
	fmt.Println(seyHello("Babi", theFilter))
}