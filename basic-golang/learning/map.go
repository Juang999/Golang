package main

import "fmt"

func main(){
	// person := map[string]string{
	// 	"name": "Bangkit Juang Raharjo",
	// 	"address": "Bogor",
	// 	"education": "Senior High School",
	// }

	person := make(map[string]string)

	person["name"] = "Bangkit Juang Raharjo"
	person["address"] = "Bogor"
	person["education"] = "Senior High School"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(len(person))
	
	delete(person, "education")
	
	person["lesson"] = "Golang"
	person["Hobby"] = "Hacking"
	
	fmt.Println(person)
	fmt.Println(len(person))
}