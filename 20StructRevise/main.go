package main

import "fmt"

type sai struct {
	name string
	roll int
	tempsai temp //now tempsai is variable's

}
type temp struct{
	godname string
	id int

}

func main() {

	see := sai{"sai", 32,temp{"krishna",23}}
	fmt.Println(see.name,see.roll,see.tempsai.godname)

	seee:=temp{"sai",1}
	fmt.Println(seee.godname,seee.id)
	fmt.Println(seee)


}
