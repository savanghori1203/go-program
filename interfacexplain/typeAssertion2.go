package interfacexplain

import "fmt"

func TypeAssertionInInterface2() {
	var i interface{} = 42

	// value := i.(string)

	switch v := i.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	case bool:
		fmt.Println("Bool", v)
	default:
		fmt.Println("Unknow type")
	}
	// fmt.Println(value)
}
