package reviesion

import "fmt"

func ArrayLearning() {
	array := [4]int{1, 2, 3, 4}
	fmt.Println(array)
	slice := array[:]
	fmt.Println(slice)

	fmt.Println(RetunArrayLearning())
}

func RetunArrayLearning() [4]int {
	var newArr [4]int

	newArr[0] = 1
	newArr[1] = 2
	newArr[2] = 3
	newArr[3] = 4

	return newArr
}
