package interfacexplain

import "fmt"

type Persion struct {
	Name string
	Age int
}

func typeAssertionWithStruct(i interface{}){

	val,ok := i.(Persion)
	if ok {
		fmt.Println("Name is " ,val.Name , " and age is " ,val.Age )
	} else{
		fmt.Println("Unkown type")
	}
}

func TypeAssertionWithStruct(){
	psr := Persion{"savan",54}
	typeAssertionWithStruct(psr)
	typeAssertionWithStruct(42)
}