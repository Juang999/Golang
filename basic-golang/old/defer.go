package main

import "fmt"

func logging(){
	message := recover()

	if message != nil {
		fmt.Println("got an error:", message)
	}

	fmt.Println("selesai menjalankan fungsi")
}

func runApp(value int){
	defer logging()

	result := 10/value
	fmt.Println("result:", result)
	fmt.Println("berhasil menjalankan App")
}

func main(){
	runApp(2)
}