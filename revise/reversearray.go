package revise

import (
	"fmt"
)

func ReverseArray(arr1 [4]int) [4]int {
	// var out [4]int
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for index, value := range arr1 {
		fmt.Println(value, index)
	}

	fmt.Println(len(arr1))

	// for i := 0; i < len(arr1); i++ {
	// 	out[len(arr1)-1-i] = arr1[i]
	// }


	//optimize solution

	n := len(arr1)

	for i := 0 ; i < n/2 ; i++{
		arr1[i], arr1[n-1-i] = arr1[n-1-i], arr1[i]
	}
	return arr1
}

func Reverse() {
	arr := [4]int{1, 2, 3, 4}
	result := ReverseArray(arr)
	fmt.Println(result)
}
