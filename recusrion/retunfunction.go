package recusrion

import "fmt"

var number int
func ReturnAnonymouseFunc(num int) func() int {
	return func() int {
		num++
		return num
	}
}

func Output(){
	x := ReturnAnonymouseFunc(number)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

/**
 return anonymouse function from the function
*/