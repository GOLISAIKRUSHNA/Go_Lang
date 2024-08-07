package main

import "fmt"

func main() {

	var arr []string
	arr1 := []string{"one","two","three"}
	fmt.Println(arr1, arr)

	arr2:=[]string{"s","u","b"}

	z:=""

	for j:= range arr2{//only index of arr2
		z+=arr2[j]
		fmt.Println(j)
	}
	fmt.Println(z)

}
