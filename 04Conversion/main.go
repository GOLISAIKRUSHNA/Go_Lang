package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("***********String to float**************")

	read := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating's for pizza:")
	input, _ := read.ReadString('\n')
	fmt.Print(input)
	fmt.Printf("type is: %T", input)
	fmt.Println("")

	hi, err:= strconv.ParseFloat(strings.TrimSpace(input),64)
	
	// fmt.Println(hi)
	// fmt.Println(err)


	// if err != nil {
	// 	fmt.Println("===========>",err)
	// } 
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(hi)
	}

}
