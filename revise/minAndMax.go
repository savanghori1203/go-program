package revise

import "fmt"

func MinAndMax(arr1 [5]int) (minNum int , maxNum int) {
	// minNum, maxNum = arr1[0], arr1[0]

	// for i := 0; i < len(arr1); i++ {
	// 	if maxNum < arr1[i] {
	// 		maxNum = arr1[i]
	// 	}
	// 	if minNum > arr1[i] {
	// 		minNum = arr1[i]
	// 	}
	// }
	// return minNum, maxNum

	//optimize solution

	if(len(arr1) == 0){
		return
	}

	minNum, maxNum = arr1[0], arr1[0]

	for _ , values := range arr1{
		if minNum > values {
			 minNum = values
		}
		if maxNum < values {
			 maxNum = values
		}
	}

	return minNum, maxNum
}

func Minandmax() {
	// arr := [5]int{76, 4, 76, 2, 89}
	var arr [5]int
	minNum , maxNum := MinAndMax(arr)
	fmt.Println(minNum)
	fmt.Println(maxNum)
}
