package main

import (
	"fmt"
	"time"
)
func main()  {
	fmt.Println("Present date and time")

	take:=time.Now()
	fmt.Println(take)
	fmt.Println(take.Format("02-01-2006 15:04:03 Monday"))



	
}