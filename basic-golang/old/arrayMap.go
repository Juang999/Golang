package main

import "fmt"

func main(){
	var array_map [5]map[string]string

	array_map[0] = map[string]string{
		"name": "Bangkit Juang Raharjo",
		"consulate": "Bogor",
	}
	array_map[1] = map[string]string{
		"name": "Faiz Marketer",
		"consulate": "Magelang",
	}
	array_map[2] = map[string]string{
		"name": "Tabir Multimedia",
		"consulate": "Bekasi",
	}
	array_map[3] = map[string]string{
		"name": "Bagus Sonarangga Ramadhan",
		"consulate": "Garut",
	}
	array_map[4] = map[string]string{
		"name": "Ismail Nuralam",
		"consulate": "Magelang",
	}

	fmt.Println(array_map[1]["name"])
}