package mapexplain

import "fmt"

func Count(){
	count := make(map[string]int)
	lang := []string{"go", "javascript", "python", "go", "javascript"}
	fmt.Println(count)

	for _,value := range lang {
		count[value]++
	}
	fmt.Println(count)

}

/**
* Here we have make object that initial value is of key is zero
* if we increment than it will increase 
*/