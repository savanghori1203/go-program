package arrayexplain

import (
	"fmt"
)

func Retunvalue() [3]int{
	return [3]int{1,2,3}
}

func Tree(){
	arr := Retunvalue()
	fmt.Println(arr)
}