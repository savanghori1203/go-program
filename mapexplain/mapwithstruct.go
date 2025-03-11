package mapexplain

import "fmt"

type Detail struct{
	Name string
	Age int
}

func MapWithStruct(){

	newMap := map[string]Detail{
		"savan":{"savan",30},
		"tirth":{"tirth",25},
	}

	fmt.Println(newMap)
}