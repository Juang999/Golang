package main

import "fmt"

func helloWorld() interface{} {
	return true
}

func main() {
	fmt.Println(helloWorld())
}