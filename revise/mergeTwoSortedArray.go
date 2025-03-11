package revise

import "fmt" 

func MergeTwoSortedArray(arr1 []int, arr2 []int) []int {
    //  for _, num := range arr2 {
		// arr1 := append(arr1,arr2...)
	//  }

	arr1 = append(arr1, arr2...)
	return arr1
}

func TwoArrayMerge(){
	arr1 := []int{4,5,6,7,8}
	arr2 := []int{7,8,9,10,11,12}
	result := MergeTwoSortedArray(arr1,arr2)
	fmt.Println(result)
}

