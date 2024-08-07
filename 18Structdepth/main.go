package main

import "fmt"

type gasengine struct {
	mpg     int
	gallons int
}


func (e gasengine) mile()int{
	return e.gallons*e.mpg

}

func makeit(e gasengine,miles int){
		 
		if miles<=e.mile(){
			fmt.Println("you make it")
		}else{
			fmt.Println("Need fuel")
		}

		
	}




func main() {

	var again gasengine = gasengine{25, 15}
	fmt.Println("mile ans",again.mile())
	var againn=struct{
		mpg int
		gallons int
	}{10,20}

	


	fmt.Println(againn.mpg,againn.gallons)
	fmt.Println(again.mpg)
	fmt.Println(again.gallons)

	fmt.Println(again.mpg,again.gallons)

	makeit(gasengine{1,2},3)


	 

}
