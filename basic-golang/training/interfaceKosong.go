package main

import "fmt"

func ups(params interface{}) interface{} {
	return params
}

func main() {
	say := ups(false)
	fmt.Println(say)
}