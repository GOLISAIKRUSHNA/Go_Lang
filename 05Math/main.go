package main

import (
	"fmt"
	"math"
	
)

func main() {
	fmt.Println("math function in go", math.Sqrt(4))
	fmt.Println(math.Ceil(2.33), math.Ceil(2.1), math.Pow(2, 4))
	fmt.Println(math.Ceil(2.33), math.Ceil(2.1), math.Round(2.5))

	fmt.Println("math function in go", math.Pi)
	fmt.Println("E:", math.E)

	fmt.Println("math function in go", math.Sqrt(4))
	degree := 30.0
	rada := degree * (math.Pi / 180)
	val := math.Sin(rada)
	fmt.Println("val of sin 30 degree:", val)

}
