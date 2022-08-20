package main

import "fmt"

func main(){
	mounth := [12]string{
		"januari",
		"vebruari",
		"maret",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	slice1 := mounth[4:7]

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}