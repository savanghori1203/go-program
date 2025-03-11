package structexplain

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.", p.Name, p.Age)
}

func One() {
	p := Person{Name: "savan", Age : 30}
	fmt.Println(p.Greet())
}
