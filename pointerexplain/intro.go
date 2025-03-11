package pointerexplain

import "fmt"

func Introduction(){
	var x int = 67

	fmt.Println(x)
	fmt.Println(&x)

	var name string = "savan"
	var ptr2 *string = &name

	fmt.Println(ptr2)
}

/**
we have to create the pointer that type variable is same as the type of variable
*/