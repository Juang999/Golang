package main

import (
	"fmt"
	"strings"
)

func main(){
	var_array := []string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jum'at",
		"sabtu",
		"ahad",
	}
	fmt.Println(strings.Join(var_array, ", "))

	slice := var_array[2:5]
	slice2 := append(slice, "minggu lagi")

	fmt.Println(slice[2])
	fmt.Println(slice)
	fmt.Println(slice2)

	len_slice := len(slice)
	fmt.Println("Panjang slice adalah:", len_slice)

	cap_slice := cap(slice)
	fmt.Println("kapasitas slice adalah:", cap_slice)

	copySlice := make([]string,  len_slice, cap_slice)
	copy(copySlice, slice)

	fmt.Println(copySlice)
}