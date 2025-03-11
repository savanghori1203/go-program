package sliceexplain

import "fmt"

func Slicing(){
	slice := []int{1,2,3,4,5,6,7,8,9}

	newSlice := slice[1:4]
	fmt.Println(newSlice)
}