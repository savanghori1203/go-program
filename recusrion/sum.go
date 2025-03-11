package recusrion

import "fmt"

func Addition(num int) int {
	if num == 0 {
		return 0
	}else{
		return num + Addition(num-1)
	}
}

func AllNumAddition(){
	x := 100
	y := Addition(x)
	fmt.Println(y)
}