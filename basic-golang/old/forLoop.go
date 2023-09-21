package main

import "fmt"

func main(){
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("perulangan ke", counter)
	// 	counter++
	// }

	// for i := 0; i <= 10; i++ {
	// 	fmt.Println("perulangan ke", i)
	// }

	var_array := [3]string{
		"Bangkit Juang Raharjo",
		"Ismail Nuralam",
		"Hayatun Nufus",
	}

	for _, value := range var_array {
		fmt.Println("value:", value)
	}
}