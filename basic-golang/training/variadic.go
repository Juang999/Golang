package main

import "fmt"

func varags(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := varags(10, 20, 30, 40, 50)
	fmt.Println(total)
}