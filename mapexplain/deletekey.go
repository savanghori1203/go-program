package mapexplain

import "fmt"

func DeleteKey(){
	newMap := map[string]string{
		"Add" : "surat",
		"street" : "3",
	}

	fmt.Println(newMap)
	delete(newMap,"Add")
	fmt.Println(newMap)
}