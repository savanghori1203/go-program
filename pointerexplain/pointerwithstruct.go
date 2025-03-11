package pointerexplain

import "fmt"

type Person struct {
	Name string
	Age  int
}

func PointerwithStruct() {
	persion1 := Person{"savan", 1}

	fmt.Println(&persion1)

	//initialize the Person type pointer which is store the value of persion varibale of struct Persion 
	ptr := &persion1
	fmt.Println(ptr.Name)
	fmt.Println(ptr.Age)
	fmt.Println(*ptr)

	//update the value using the persion pointer 
	ptr.Name = "safi"
	ptr.Age = 23

	fmt.Println(persion1) //new value will be printed 
}
