package main

import "fmt"

func main() {

	// var arrname [3]int

	var arrname =[...]string{"goli","sai","anil","sameer"}
	fmt.Println(arrname)

	var apna= map[string]int{"goli":1,"sai":2,"krishna":3}
	// var hi,sai=apna["sai"]

	// if sai{
	// 	fmt.Println(hi)
	// }

	// fmt.Println(apna)

// 	var i int=0

// 	for i<10{
// 	fmt.Println(i)
// 	i=i+1
// }

// for j:=0;j<10;j++{
// 	fmt.Println(j)
// }



// for name,age :=range apna{
// 	fmt.Println(name,age)
// }

// for x,y:=range apna{
// 	fmt.Println(x,y)
// }

// for i,j:= range arrname{
// 	fmt.Println(i,j)
// }



for i,j:=range apna {
	fmt.Println(i,j)

}
}
