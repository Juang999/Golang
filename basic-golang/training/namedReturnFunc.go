package main

import "fmt"

func callName() (fullName, nickName, job string){
	fullName = "Bangkit Juang Raharjo"
	nickName = "Juang"
	job = "Backend Developer"

	return
}

func main(){
	fullName, nickName, job := callName()
	fmt.Println(fullName)
	fmt.Println(nickName)
	fmt.Println(job)
}