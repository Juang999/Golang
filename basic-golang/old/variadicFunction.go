package main

import "fmt"

func sumAll(data ...int) int {
	total := 0

	for _, value := range data {
		total += value
	}

	return total
}

func main(){
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)

	thisSlice := []int{10, 10, 10, 10, 10, 10}
	total1 := sumAll(thisSlice...)
	fmt.Println("ini adalah data dari slice")
	fmt.Println(total1)
}