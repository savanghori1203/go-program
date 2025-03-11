package sliceexplain

import "fmt"

func ClearSlice(){
  slice := []int{1,2,3,4,5,6,7,8}
  slice = slice[:0]
  fmt.Println(slice)	
}