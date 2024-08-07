package main

import (
	"fmt"

	
)

func main() {
	fmt.Println("Map in Depth")

	var mapdepth= map[string]int{"Apple ":1,"banana":2}

	fmt.Println("chalooo",mapdepth)

	for i,j:=range mapdepth{
		fmt.Println(i,j)
	}

	 mapdepth2:= map[string]int{"Apple ":1,"banana":2,"Orange":3}
	 fmt.Println("volvo:",mapdepth2)

	 delete(mapdepth2,"banana")

	 fmt.Println("volvo:",mapdepth2)


	 fmt.Println("*************Make Function**************")

	 var see=make(map[string]int) //with help of make()

	 var see1 map[int]string

	 

	 fmt.Println(see==nil)
	 fmt.Println(see1==nil)

	 check,key:= mapdepth2["Apple"]

	 fmt.Println(check,key)

var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}
fmt.Println(a)
	  var b= map[string]string{"one":"first","second":"second"}
	  fmt.Println(b)

	  ok,hi:=b["ne"]
	  fmt.Println(ok ,hi)

	  
// val1, ok1 := a["brand"] 
// 	  fmt.Println(val1,ok1)




var z=map[string]string{"one":"first","sec":"second bro"}
x:=z


fmt.Println(z)

x["one"]="value"
fmt.Println(x)
x["sec"]="sab change"
fmt.Println(z)
fmt.Println(x)


var xyz[]string

ab:=append(xyz,"one","two","three")

fmt.Println(ab)

for aja,j:=range ab{
	fmt.Println(aja,j)
}

var arr=[3]int{1,2,3}



var arr1=[...]int{1,2,3}
fmt.Println(arr1)
fmt.Println(arr)


var arr3=[3]int{1,2,3}
fmt.Println(arr3)

var ar5=[3]int{4,5,6}
fmt.Println(ar5)
arr4:=[...]int{1,2,3}
arr5:=[3]int{1,2,3}
fmt.Println(arr4)
fmt.Println(arr5)



}
