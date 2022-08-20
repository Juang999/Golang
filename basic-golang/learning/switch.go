package main

import "fmt"

func main(){
	name := "Mahameru"

	switch name {
	case "mahmud":
		fmt.Println("hello ", name)
	case "Bangkit Juang Raharjo":
		fmt.Println("hello ", name)
	default:
		fmt.Println("ga kenal ahhhhhh")
	}
}