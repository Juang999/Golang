package main

import "fmt"

func main(){
	name := "Juang"
	counter := 0

	increment := func() {
		name = "Mujahid"
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}