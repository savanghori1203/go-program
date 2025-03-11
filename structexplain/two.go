package structexplain 

import "fmt"

type Address struct {
	Name string
	Sreet string
	City string
	Pincode int
}

func Test(){
	address := Address{Name : "savan", Sreet : "test", City : "surat", Pincode : 1234567}
	fmt.Printf("Address %+v\n", address)
	address.Greet()
}

func (p Address) Greet() {
	savan := Address{Name : "Tirth", Sreet : "test", City : "junagadh", Pincode : 876543}
	fmt.Println(savan)
}

