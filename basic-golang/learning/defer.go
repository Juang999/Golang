package main

import "fmt"

func logging() {
	fmt.Println("selssai memanggil function")
}

func runApp() {
	defer logging()
	fmt.Println("Run the App")
}

func main() {
	runApp()
}