package main

import "fmt"

func main(){
	map_go := map[string]string{
		"first_key": "first_value",
		"second_key": "second_value",
		"third_key": "third_value",
	}
	
	fmt.Println(map_go)
	fmt.Println(map_go["first_key"])
	fmt.Println(map_go["second_key"])
	fmt.Println(map_go["third_key"])

	map_go["fourth_key"] = "this is the different"
	fmt.Println(map_go["fourth_key"])
}