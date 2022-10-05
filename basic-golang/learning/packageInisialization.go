package main

import (
	"Golang/basic-golang/learning/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}