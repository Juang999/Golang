package main

import "fmt"

func numbers(limit ...int) (int, int) {
	total := 0
	numbers := 0

	for _, realNumber := range limit {
		total += realNumber
		numbers++
	}

	return total, numbers
}

func main(){
	total, list := numbers(1,2,3,4,5)
	fmt.Println(total, list)
}