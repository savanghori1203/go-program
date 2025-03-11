package sliceexplain

import "fmt"

func Iterating(){
	slice := []int{1,2,3,4,5,6,7,8}

	for i,v := range slice {
		fmt.Println(i,v)
	}

	for i := 0; i< len(slice) ; i++ {
		fmt.Println(slice[i])
	}
}