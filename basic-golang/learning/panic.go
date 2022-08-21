package main

import "fmt"

func endApp() {
	catch := recover()
	if catch != nil {
		fmt.Println("ERROR message;", catch)
	}
	fmt.Println("Applikasi selesai")
}

func runApp1(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp1(true)
}