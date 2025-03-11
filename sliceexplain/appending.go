package sliceexplain

import "fmt"

func Appending() {

	//methode 1
	slice := make([]int, 3)

	slice[0] = 1
	slice[1] = 2
	slice[2] = 3

	fmt.Println(slice)

	//mehode 2
	var slice2 []int

	slice2 = append(slice2, 1)
	slice2 = append(slice2, 4)
	slice2 = append(slice2, 98)

	fmt.Println(slice2)

	/**
	* when you want to use slice[0] methode to add value in slice we need to allocate the memory
	before the it will be use by using the make([]int,size) using this need to allocate the memory
	*/

	/**
	* if you use var slice2 []int to initialize the slice now if you not allocate the memory
	* you need to use append() methode for add in slice
	 */

	/**
	* In append(destinaton,source) in this we have to add data in "destination"
	*/

	slice3 := []int{1,2,3,4}
	slice3 = append(slice3, 3,5)

	fmt.Println(slice3)
}
