package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println("Error handling in go")

	a,b,err:=sai(5,0)

	// if err!=nil{
	// 	fmt.Println(err.Error())
	// }else if b==0{
	// 	fmt.Println("remainder is zero")
	// }else{
	// 	fmt.Printf("%d %d",a,b)
	// }


	// switch{
	// 	case err!=nil:{
	// 		fmt.Println(err.Error())
	// 	}
	// case b==0:
	// 	fmt.Println("remainder of b is:",b)
	// default:
	// 	fmt.Println("no error")
	// }
	// fmt.Println(a,b)


	switch err{
	case err:
		fmt.Println("error in code")

		fmt.Println(err.Error())
		
	
	default:
		fmt.Println(a,b)
		fmt.Println("error in code")
	}


	
}

func sai(x int ,y int) (int,int,error){

	var see error
	
	if y==0{

		aya:=errors.New("zero division error")

		return 0,0,aya
		


	}


	// fmt.Println(see)

	val1:=x/y
	val2:=x%y
	
	return val1,val2,see

}
