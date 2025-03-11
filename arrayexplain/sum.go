package arrayexplain

import "fmt"

func Sum() {
	var sum int
	arr := [5]int{54, 23, 1, 9, 78}
	for _, value := range arr {
      sum += value
	}

	fmt.Println(sum)
}
