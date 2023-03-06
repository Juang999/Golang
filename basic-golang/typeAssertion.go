package main

import (
	"fmt"
	"strings"
)

func main() {
	thisString := []string{"Hello", "World"}

	fmt.Println(strings.Join(thisString, "-"))
}