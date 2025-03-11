package reviesion

import "fmt"

var Fullname string

func VariableLearing() {

	name := "savan"

	fmt.Println(name)

	var surname = "ghori"
	fmt.Println(surname)

	Fullname = name + " " + surname
	fmt.Println(Fullname)

	age := 20
	fmt.Println(age)

	const price = 100
	fmt.Println(price)

	a := 10
	b := 20
	fmt.Println(ReturnValue(a, b))
}

func ReturnValue(a int, b int) int {
	return a + b
}
