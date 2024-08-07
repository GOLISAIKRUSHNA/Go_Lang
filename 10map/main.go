package main

import (
	"fmt"
)

func main() {


	// mapping in go 

	// var mapname map[string]int=make(map[string]int)
	// fmt.Println(mapname)

	
	// var mapsecond map[string]int=make(map[string]int)
	// fmt.Println(mapsecond)
	var mapsecond2= map[string]int{"anil":1,"tejas":2,"bhavya":3}
	fmt.Println(mapsecond2["bhavya"])

	var age,ok=mapsecond2["bhavya"]
	if ok{
		fmt.Println("age value is:",age)
	}else{
		fmt.Println("not there in map")
	}


	var try=map[string]int{"first":1,"sec":2}
	fmt.Println(try)


	var arr=[...]int{1,2,3}
	fmt.Println(arr)

	arrrr:=[3]int{4,5,6}
	fmt.Println(arrrr)


	for i,j:=range arrrr{
		fmt.Println(i,j)
	}

	for i:=0;i<len(arrrr);i++{
		fmt.Println(arrrr[i])
	}
	

}
