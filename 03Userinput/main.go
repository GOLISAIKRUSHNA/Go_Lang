package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("input")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the user input:")

	input, _:= reader.ReadString('\n')
	// fmt.Println(input)
	fmt.Printf("Rated for pizza : %T", input)
	fmt.Println("")

	
}
