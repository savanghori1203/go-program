package pointerexplain

import "fmt"

func Pointer3() {
	num := 9876
	fmt.Println(num)

	takePointer(&num)

	fmt.Println(num)
}

func takePointer(ptr *int) {
	*ptr = 34
}
