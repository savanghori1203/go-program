package arrayexplain

import (
	"fmt"
	"sort"
)

func FindSecondLargest(arr [4]int) int {
  slice := arr[:]  
  sortArray := sort.IntSlice(slice)
  fmt.Println(sortArray)
  return arr[1]
} 

func Largest(){
    arr := [4]int{87,43,7,8}
    result := FindSecondLargest(arr)
    fmt.Printf("Second largest number is:  %d\n",result)
}