package main

import (
	"os"
	"fmt"
)

func main(){
	args := os.Args
	fmt.Println(args)

	name, err := os.Hostname()

	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println("error: ", err.Error())
	}
}