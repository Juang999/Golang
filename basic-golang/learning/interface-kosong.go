package main

import "fmt"

func returnAny(anyParameter interface{}) interface{} {
	return anyParameter
}

func main() {
	var name interface{} = returnAny(4)
	fmt.Println(name)
}