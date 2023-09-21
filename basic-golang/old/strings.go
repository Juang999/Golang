package main

import (
	"fmt"
	"strings"
)

func main(){
	fullName := []string{
		"Bangkit",
		"Juang",
		"Raharjo",
	}

	concat := strings.Join(fullName, " ")

	fmt.Println(concat)
}