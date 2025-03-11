package sliceexplain

import "fmt"

func DeleteElement(){
	slice := []int{1,2,3,4,5,6,7,8,9}
	index := 3

	slice = append(slice[:index], slice[index+1:]...)
	fmt.Print(slice)
}