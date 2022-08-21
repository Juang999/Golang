package main

import "fmt"

type BlacklistWord func(string) bool

func registerUser(name string, blacklistWord BlacklistWord) string {
	if blacklistWord(name) {
		return "hello " + name
	} else {
		return "name " + name + " has been blocked"
	}
}

func main(){
	blacklist := func(name string) bool {
		return name != "anjing"
	}

	fmt.Println(registerUser("anjing", blacklist))
}