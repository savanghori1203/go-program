package revise

import "fmt"

func SumOfArray(arr [7]int) (result int) {
   for _,value := range arr {
	 result += value
   }
   return result
}

func ArraySum(){
	arr := [7]int{98,3,56,4,32,4,123}
	result := SumOfArray(arr)
	fmt.Println(result)
}