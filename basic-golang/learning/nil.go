package main

import "fmt"

func whatsUp(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	call := whatsUp("Bangkit Juang Raharjo")

	fmt.Println(call)
}