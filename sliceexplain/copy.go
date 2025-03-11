package sliceexplain

import "fmt"

func Copy(){
	slice := []int{7,6,5,4,3,2,1}
	destination := make([]int,len(slice))
	copy(destination,slice)

	slice[0] = 98
	fmt.Println(destination)
	fmt.Println(slice)
}