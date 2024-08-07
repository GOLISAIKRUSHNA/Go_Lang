package main

import "fmt"

type apna struct {
	mpg         int
	gallon      int
	ownerinfo   owner
	rentinfo    renter
	studentinfo student
	own
	int
}
type own struct{
	name string
}
type student struct {
	name    string
	rollno  int
	college string
}

type owner struct {
	name string
}
type renter struct {
	namerent string
}

func main() {
	var second apna = apna{1, 2, owner{"ramp.ai solution"}, renter{"Saikrushna"}, student{"sai", 32, "atharva college"},own{"apna college"},3}
	fmt.Println(second.gallon)
	second.gallon = 3
	fmt.Println(second.gallon, second.mpg, second.ownerinfo.name)
	fmt.Println(second.rentinfo.namerent, second.ownerinfo.name)
	fmt.Println(second.studentinfo.name, second.studentinfo.college, second.studentinfo.rollno)
	fmt.Println(second.own.name,second.int)

}
