package pointerexplain

import "fmt"

func Callbyvalue(y int) {
	y = 30
	fmt.Println(y)
}

func CallbyRefrence(x *int){
   *x = 87
   fmt.Println(x)
}

func CallBy(){
	var num int

	Callbyvalue(num)
	fmt.Println(num)

	CallbyRefrence(&num)
	fmt.Println(num)
}