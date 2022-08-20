package main

import "fmt"

func main(){
	// loop for map data
	// map_variable := map[string]string{
	// 	"name": "Bangkit Juang Raharjo",
	// 	"age": "23",
	// 	"job": "Backend Programmer",
	// }

	// fmt.Println(map_variable)

	// for index, value := range map_variable {
	// 	fmt.Println("index", index, "adalah", value)
	// }

	// loop for array
	// array_variable := [...]string{
	// 	"Bangkit Juang Raharjo",
	// 	"23",
	// 	"Backend Programmer",
	// }

	// fmt.Println(array_variable)

	// number := 0

	// for number <= 2 {
	// 	fmt.Println(array_variable[number])

	// 	number++
	// }

	// for loop together
	array_variable := [...]string{
		"Bangkit Juang Raharjo",
		"23",
		"Backend Programmer",
	}

	for i := 0; i < len(array_variable); i++ {
		fmt.Println("array ke-",i, "memiliki value:",array_variable[i])
	}
}