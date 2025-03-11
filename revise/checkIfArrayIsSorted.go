package revise

import "fmt"

func CheckIfArrayIsSorted(arr [6]int) bool{
 for i := 1 ; i < len(arr) ; i++ {
	if arr[i] < arr[i-1]{
		return false
	}
 }
 return true
}

func CheckArraySoretd(){
	arr := [6]int{1,2,3,4,5,6}
	result := CheckIfArrayIsSorted(arr)
	fmt.Println(result)
}