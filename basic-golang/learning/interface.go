package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

// func (person Person) getName

func main() {
	var juang Person

	juang.Name = "Bangkit Juang Raharjo"
	sayHello(juang)

}