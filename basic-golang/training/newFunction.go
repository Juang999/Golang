package main

import "fmt"

func callHisName(fullName string, numbers ...int) (name string, total int) {
	name = fullName
	total = 0

	for _, value := range numbers {
		total += value
	}

	return
}

func main(){
	name, number := callHisName("Bangkit Juang Raharjo", 10, 20, 30)
	fmt.Println(name)
	fmt.Println(number)
}