package stringexplain

import (
	"fmt"
	"strings"
)

func StringExplain() {
	task := "Have to complete all the string program"

	// for i :=0 ; i < len(task) ; i ++ {
	// 	fmt.Println(task[i])
	// }

	fmt.Printf("%c\n", task[0])
	fmt.Println(len(task))

	//join two string
	first := "Hello"
	second := "savan"

	thirid := first + " " + second
	fmt.Println(thirid)

	result := strings.Compare(first, second)
	fmt.Println(result)

	result2 := strings.Contains(task, "to")
	fmt.Println(result2)

	result3 := strings.Replace(task, "to", "the", 1)
	fmt.Print(result3)

    result4 := strings.Split(task, "")
	fmt.Println(result4)

	//screte the string from the slice

	newSlice := []string{"i", "love", "Programming"}

	result5 := strings.Join(newSlice, " ")
	fmt.Println(result5)

}
