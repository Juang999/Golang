package main

import "fmt"

func main(){
	// name := "Raharjo"

	// switch name {
	// case "Bangkit":
	// 	fmt.Println("Hello", name)
	// case "Juang":
	// 	fmt.Println("Hello", name)
	// default:
	// 	fmt.Println("Hello, boleh kenalan?")
	// }

	switch name := "Raharjo"; name {
	case "Bangkit Juang Raharjo":
		fmt.Println("hello", name)
	case "Fulan Bin Fulan":
		fmt.Println("hello", name)
	default:
		fmt.Println("apasih kamu, jomblo itu ngga diajak tau ga?")
	}
}