package main

import "fmt"

func main() {
	name := "Bangkit Juang Raharjo"
	counter := 0
	increment := func() {
		name := "Mujahid"
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}