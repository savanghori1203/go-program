package sliceexplain

import "fmt"

func edit(slice []int) []int {
	newSlice := slice[:3]
	fmt.Println(newSlice)
	slice[5] = 100
	return newSlice
}

func PassedInFunction() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	newSlice := edit(slice)
	fmt.Println(newSlice)
	fmt.Println(slice)
}
