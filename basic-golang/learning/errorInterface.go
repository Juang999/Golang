package main

import (
	"errors"
	"fmt"
)

func pembagi(bilangan int, pembagi int) (int, error) {
	if pembagi == 0 {
		return pembagi, errors.New('bilangan pembagi adalah', pembagi)
	} else {
		return bilangan / pembagi, nil
	}
}

func main() {
	var contohError error = errors.New("ini error")
	fmt.Println(contohError)
}