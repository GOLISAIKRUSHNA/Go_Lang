package main

import "fmt"

func main(){
	fmt.Println("Recursive call")
	recur(1)

}

func recur(x int) int{
	if x==10{
		fmt.Println("bas kar 10 element ayaaa")
		return 0
	}
	
	
	fmt.Println(x)
	
	return recur(x+1)





}

// func main() {

// 	fmt.Println("program starts")
// 	sai(3)
// 	res,result:=add_rem(2,4)

// 	fmt.Println(res,result)


// }


// func sai(x int)int {
// 	fmt.Printf("User defined function: %d \n",x)
// 	hi("goli")
// 	hi("welcome")
// 	hi("to fintech")
// 	h:=cal(3.5)
// 	fmt.Println(h)

// 	return x

// }

// func hi(x string)string {
// 	fmt.Printf("User defined function: %s \n",x)
// 	fmt.Println(x)
// 	return x

// }

// func cal(x float64 ) float64{
// 	// fmt.Println(x)

// 	return x

// }

// func add_rem(x int,y int ) (int,int){
// 	// fmt.Println(x)

//      val1:=x
//      val2:=y

// 	 return val1,val2


// }

