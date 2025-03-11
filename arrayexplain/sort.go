package arrayexplain

import (
	"fmt"
	"sort"
)

func Methode(){
  arr := [...]int{9,3,5,2,8,1}
  fmt.Println(arr)

  slice := arr[:]
  sort.Sort(sort.Reverse(sort.IntSlice(slice)))

  fmt.Println(slice)
}
