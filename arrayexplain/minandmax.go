package arrayexplain

import (
	"fmt"
)

func MinMax(){
	arr := [5]int{8,13,65,23,78}

	min := arr[0]
	max := arr[0]
	for _,value := range arr {
       if value > max {
		 max = value 
	   }
	   if value < min {
		 min = value
	   }
	}

	fmt.Println(min,max)
}