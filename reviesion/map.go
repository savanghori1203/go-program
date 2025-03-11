package reviesion

import "fmt"

func MapLearing() {
	newMap := map[string]int{"age": 23, "exp": 1}
	fmt.Println(newMap)

	for index, value := range newMap {
		fmt.Println(index)
		fmt.Println(value)
	}

	 obj := ReturnMap()
	 fmt.Println(obj)
}

func ReturnMap() map[string]int {
	return map[string]int{"count": 3, "after": 6}
}
