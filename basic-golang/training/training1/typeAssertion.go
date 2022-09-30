package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	var result interface{} = random()
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown")
	} 
}