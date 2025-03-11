package pointerexplain

import "fmt"

func Pointer(){
	var ptr = new(int)

	fmt.Println(ptr) //zero value
	fmt.Println(*ptr) //address of that pointer

	*ptr = 34 //assingn the value of that pointer

	fmt.Println(ptr) //address
	fmt.Println(*ptr) //34
}


/**
Here starting pointer we initialize the pointer and value is zero becuase it is interger type
*/
