package main

import "fmt"

type Blacklist func(string) bool

func register(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked! ",name)
	} else {
		fmt.Println("welcome ",name)
	}
}

func blacklist(name string) bool {
	
}

func main() {

}