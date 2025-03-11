package revise

import (
	"fmt" 
	"sort"
)

func RemoveDuplicateFromArray(arr [6]int) []int {
	slice := arr[:]
	newSlice := []int{}
	uniqueMap := make(map[int]bool)
    for _,value := range slice {
		if(!uniqueMap[value]){
			uniqueMap[value] = true
			newSlice = append(newSlice, value)
		}
	}

	sort.Ints(newSlice)
	return newSlice
}

func DuplicateRemove(){
	arr := [6]int{5,3,4,5,3,6}

	result := RemoveDuplicateFromArray(arr)
	fmt.Println(result)
}