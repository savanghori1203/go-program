package reviesion

import "fmt"

type Persion struct {
	Name string
	Age  int
	Exp  int
}

func (p Persion) Greet() {
	fmt.Println(p.Name,p.Age,p.Exp)
}

func StructLearning() {
	obj := Persion{Name: "savan", Age: 23, Exp: 3}
	obj.Greet()
	fmt.Println(obj)
	obj.Welcome()
}


func (p Persion) Welcome(){
	savan := Persion{Name: "tirth", Age: 23, Exp: 6}
	fmt.Printf("Address %+v\n",savan)
	fmt.Println(p.Name)
}