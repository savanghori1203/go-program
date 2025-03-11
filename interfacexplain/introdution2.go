package interfacexplain

import (
	"fmt"
	"math"
)

type Shape interface{
	Area() float32
}

type Circel struct{
	Radius float32
}

func (c Circel) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct{
	Width float32
	Hight float32
}

func (r Rectangle) Area() float32{
	return r.Hight * r.Width
}

func PrintArea(s Shape){
	fmt.Println(s.Area())
}

func ShapeArea(){
	c := Circel{5.10}
	r := Rectangle{5.1, 8.5}

	PrintArea(c)
	PrintArea(r)
}