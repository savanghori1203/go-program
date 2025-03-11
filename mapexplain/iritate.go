package mapexplain

import "fmt"

func Iritate(){
	fruits := map[string]int{"apple": 10, "banana": 5, "orange": 8}

	for key,value := range fruits{
		fmt.Println(key,value)
	}

}