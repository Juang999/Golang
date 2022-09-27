package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("terjadi error:", message)
	}
	fmt.Println("applikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("applikasi error")
	}

	fmt.Println("applikasi berjalan")
}

func main() {
	runApp(true)
}