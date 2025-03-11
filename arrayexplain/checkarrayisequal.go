package arrayexplain

import (
	"fmt"
)

func Checkisequal(arr1,arr2 [3]int) bool {
   for i := 0 ; i < len(arr1) ; i++ {
	 if arr1[i] != arr2[i] {
		return false
	 }
   }
   return true
}

func Testisequal() {
	arr1 := [3]int{1,2,3}
	arr2 := [3]int{1,2,3}
	result := Checkisequal(arr1,arr2)
	fmt.Println("is both array is equal: " , result)
}