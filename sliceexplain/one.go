package sliceexplain

import "fmt"

func Slice() {
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	slice := arr[1:5]
	fmt.Println(slice)

	for _,value := range arr {
		fmt.Println(value)
	}
}
