package interfacexplain

import "fmt"

type Animal struct {
	Name string
}

func TypeAssertionWithPointer(){
	c := Animal{"Dog"}
	fmt.Println(c)
	fmt.Println(&c)
	ptr := &c
	fmt.Println(*ptr)

	var i interface{} = &c
	a,ok := i.(*Animal)
	if ok {
		fmt.Println("Animal", a.Name)	
	}else {
		fmt.Println("Not an animal")
	}
}