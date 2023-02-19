package main

import "fmt"

func main(){
	// for counter := 0; counter <= 10; counter++ {
	// 	if counter == 5 {
	// 		break
	// 	}
	// 	fmt.Println("perulangan ke", counter)
	// }

	for counter := 0; counter <= 10; counter++ {
		if counter%2 == 0 {
			continue
		}
		fmt.Println("perulangan ke", counter)
	}
}