package pointerexplain

import "fmt"

var x int = 20

func Savan(ptr *int) {
	*ptr = 20
	fmt.Println(*ptr)
}

func test() int {
	x = 10
	fmt.Println(x)
	return x
}

func Test2(){
	x = test()
	Savan(&x)
	fmt.Println(x)
}
