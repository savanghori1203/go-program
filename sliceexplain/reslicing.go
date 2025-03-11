package sliceexplain

import "fmt"

func Reslicing(){
	num := []int{9,8,7,6,5,4,3,2}
	slice := num[:2]
	fmt.Println(slice)

	neslice := num[:cap(num)]
	fmt.Println(neslice)
}