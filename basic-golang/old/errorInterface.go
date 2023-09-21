package main

import (
	"errors"
	"fmt"
)

func pemanggilanError(nilai1 int, nilai2 int) (int, error) {
	if nilai2 == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	} else {
		return nilai1/nilai2, nil
	}
}	

func main(){
	fmt.Println(pemanggilanError(10, 0))
}