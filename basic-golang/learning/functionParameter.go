package main

import "fmt"

type Filter func(string) string

func sayHelloFilter(name string, filter Filter) (greating string) {
	greating = filter(name)
	return
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	fmt.Println(sayHelloFilter("Bangkit Juang Raharjo", spamFilter))
	fmt.Println(sayHelloFilter("Anjing", spamFilter))
}