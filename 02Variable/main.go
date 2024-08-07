package main

import (
	"fmt"
	"unicode/utf8"
)

const Hi int=1


func main(){
	fmt.Println("*******Varible's********")

	var number int =45
	num:=23
	lst:=34
	fmt.Println(lst)
	fmt.Println(num)
	fmt.Println(number)
	fmt.Printf("This is: %T ",number)
	fmt.Println("")

	floatvalues := 12.333344444334543443
	floatvaluess :=13.222234543234543443454
	fmt.Println(floatvalues)
	fmt.Println(floatvaluess)
	fmt.Printf("This is: %T ",floatvalues)
	fmt.Println("")

	var stringvalues string= "apna college"
	strr:="values"
	fmt.Println(strr)
	fmt.Println(stringvalues)
	fmt.Printf("This is: %T ",stringvalues)
	fmt.Println("")

	var boolean bool= true
	fmt.Println(boolean)
	fmt.Printf("This is: %T ",boolean)
	fmt.Println("")

	var comp complex128= complex(1,2)
	mathcomplex:=complex(1,3)
	fmt.Println(mathcomplex)
	fmt.Println(comp)
	// fmt.Printf("This is: %T ",comp)
	fmt.Println("")

	// implicitly
	var str ="goli"
	// no var
	strval:="sai"
	fmt.Println(str)
	fmt.Println(strval)

	fmt.Println(Hi)


	var x int=10
	var y int=20

	z:=30
	

	
	fmt.Println("x and y variable",x,y,z)


	val:=int(floatvalues)
	fmt.Println(val)

	var asic rune='a'
	var asics rune='c'
	fmt.Println(asic)
	fmt.Println(asics)


	fmt.Println(len("hiiii"))


	fmt.Println(utf8.RuneCountInString("goli"))


	var a,b="goli","tech"
	fmt.Println(a,b)

	const c float32 = 1.23333
	fmt.Println(c)

	const set int=0
	fmt.Println(set)

	var charc rune='a'
	fmt.Println(charc)

	var character string="s"
	fmt.Println(character)

	var noint int
	fmt.Println(noint)

	var nofloat float32
	fmt.Println(nofloat)
	
	var nostring string
	fmt.Println(nostring)


	
}
