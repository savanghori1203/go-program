/**
*Write a function rotateRight that takes an array and an integer k as input and rotates the array to the right by k positions.
For example, if the array is [1, 2, 3, 4, 5] and k = 2, the output should be [4, 5, 1, 2, 3].
*/

package sliceexplain

import "fmt"

func RightShifting(){
	arr := [6]int{1,2,3,4,5,6}
	slice := arr[:]
	index := 6
	slice = append(slice[(len(slice)-index):], slice[:(len(slice)-index)]...)

	fmt.Println(slice)

}