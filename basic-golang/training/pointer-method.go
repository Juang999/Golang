package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "Mr."+ man.Name
}

func main() {
	Juang := Man{"Juang"}
	Juang.married()

	fmt.Println(Juang.Name)
}