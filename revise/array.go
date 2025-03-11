package revise

import (
	"fmt"
)

func Arrayexplain(){
	var arr [3]int

	fmt.Println(arr) // Initial value is zero

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)

	arr2 := [3]string{"savan","hiren", "ravi"}
	fmt.Println(arr2)

	//make slice from the array
	slice := arr2[:3]
	fmt.Println(slice)

	//take value from the function
	arr3 := RetunArray()
	fmt.Println(arr3)

	//pointer
	Pointer()
}

func RetunArray() [3]int {
	arr1 := [3]int{1,2,3}
	return arr1
}

func TakePointer(ptr *[3]int) {
    fmt.Println(ptr)
	fmt.Println(*ptr)
}

func Pointer(){
	arr := [3]int{1,2,3}
	fmt.Println(&arr)
	TakePointer(&arr)
}