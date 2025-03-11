package pointerexplain

import "fmt"

func Introduction2(){
	var num int
	var ptr *int


	num = 22 
	fmt.Println(num)
	fmt.Println(&num)

	ptr = &num
	fmt.Println(ptr)
	fmt.Println(*ptr)

	*ptr = 543
	fmt.Println(ptr)
	fmt.Println(*ptr)
	fmt.Println(num)
	fmt.Println(&num)
}