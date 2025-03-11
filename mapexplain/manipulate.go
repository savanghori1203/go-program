package mapexplain

import "fmt"

func Manupulate(){
	newMap := map[string]int{
		"apple" : 1,
		"banana" : 2,
	}

	fmt.Println(newMap["apple"])

	_,value := newMap["mango"]

	if value {
		fmt.Println("mango exist with the value", value)
	}else {
		fmt.Println("mango not exist")
	}
}