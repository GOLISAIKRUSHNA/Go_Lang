package main

import (
	"fmt"
)

func main() {

	fmt.Println("Array syntax in go")

	var arr [3]int

	fmt.Println(arr)
	fmt.Println(arr[0])

	fmt.Println(arr[1])

	arr[0] = 3
	arr[1] = 2

	arr[2] = 1

	fmt.Println(arr)

	newarr := [...]int{1, 2, 3}

	fmt.Println(newarr)

	seearr := [...]int{6, 7, 8}
	fmt.Println(seearr)

	fmt.Println(len(seearr), cap(seearr))

	finalsee := append(seearr[:], 9)

	fmt.Println(len(finalsee), cap(finalsee))
	fmt.Println(finalsee)

	fin := append(finalsee[:], 9)
	fmt.Println(fin, len(finalsee), cap(finalsee))

	fmt.Println(fin[:2])

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}

	naya := append(myslice2, "goli")
	fmt.Println(naya)
	fmt.Println(fin, len(naya), cap(naya)) //to bhavya

	s := append(naya, myslice2...)
	fmt.Println(s)


	//mapping in go 

	// var mapname map[string]int=make(map[string]int)
	// fmt.Println(mapname)

	
	// var mapsecond map[string]int=make(map[string]int)
	// fmt.Println(mapsecond)
	var mapsecond2= map[string]int{"anil":1,"tejas":2}
	fmt.Println(mapsecond2["anil"])


	var str = "string passed"


	fmt.Println(str)

	// for _,j:=range str{
	// 	fmt.Println(j)
	// }
	fmt.Printf(" %s,%T \n",str,str)

	fmt.Println(len(str))

}
