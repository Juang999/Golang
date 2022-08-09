package main

import "fmt"

func main(){
	// array := [...]string{
	// 	"januari",
	// 	"februari",
	// 	"maret",
	// 	"april",
	// 	"mei",
	// 	"juni",
	// 	"juli",
	// 	"agustus",
	// 	"september",
	// 	"oktober",
	// 	"november",
	// 	"desember",
	// }

	// fmt.Println(len(array))

	// for i := 0; i < len(array); i++{
	// 	fmt.Println("Bulan", array[i])
	// }

	maps := map[string]string{
		"name": "Bangkit Juang Raharjo",
		"email": "juangraharjo03@gmail.com",
		"phone": "081325507780",
	}

	for index, value := range maps {
		fmt.Println("index", index, "=", value)
	}
	
}