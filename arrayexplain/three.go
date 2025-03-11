package arrayexplain

import (
	"fmt"
)

func Four(){
	arr := [5]int{1,2,3,4,5}
	fmt.Println(arr)
	arr2 := [...]string{"Go","paython","java"}
	fmt.Println(arr2)

	for i := 0; i < len(arr) ; i++ {
		fmt.Println(arr[i])
	}

	for index,value := range arr {
		fmt.Println(index,value)
	}

	// Creates a copy, not a reference
	arr3 := arr
	fmt.Println(arr3)
	arr3[0] = 8
	fmt.Println(arr3)
	fmt.Println(arr)
}

func ArrayReturn(arr [4]string){
	arr[0] = "new"
}

func Final(){
	myArr := [4]string{"savan","hiren","tirth","nishit"}
	myArr2 := myArr

    ArrayReturn(myArr)
	fmt.Println(myArr) 
	fmt.Println(myArr == myArr2)

	slice := myArr2[:]
	fmt.Println(slice)
}

