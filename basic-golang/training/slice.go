package main

import "fmt"

func main(){
	mounths := [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	slice1 := mounths[2:6]
	slice2 := mounths[10:11]

	fmt.Println(append(slice1, "juni lagi"))
	fmt.Println(append(slice2, "desember lagi"))
	fmt.Println(mounths)
}