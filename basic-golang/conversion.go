package main

import "fmt"

func main(){
	var nilai32 int32 = 1234

	var nilai64 = int64(nilai32)
	var nilai16 = int16(nilai32)
	var string32 = string(nilai32)
	var boolean32 = bool(true)

	var name string = "Bangkit Juang Raharjo"
	var e = name[2]

	eString := string(e)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(string32)
	fmt.Println(boolean32)
	fmt.Println(eString)
} 