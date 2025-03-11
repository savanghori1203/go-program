package arrayexplain

import "fmt"

func Search(arr [5]int , number int) bool {
	for _, value := range arr {
		if value == number {
			return true
		}
	}
	return false
}

func Output(){
	arr := [5]int{8,13,65,23,78}
	out := Search(arr , 65)
	fmt.Println(out)
}