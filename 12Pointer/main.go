package main

import "fmt"

func main() {

	value:=3
	var a *int  //pointer alway's store's address
	a = &value

	fmt.Println(value)
	fmt.Println(a)

	fmt.Println(&value)




}
