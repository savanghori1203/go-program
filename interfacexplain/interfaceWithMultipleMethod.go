package interfacexplain

import "fmt"

type Vehicale interface{
	start() string
	stop() string
}

type Car struct{
	Name string
}

func (c Car) start() string{
	result := c.Name + " "+"START"
	return result
}

func (c Car) stop() string{
	result := c.Name + " " + "STOP"
	return result
}

func Status(s Vehicale){
	fmt.Println(s.start())
	fmt.Println(s.stop())
}

func CarDriving(){
	c := Car{"harrier"}
	Status(c)

	obj := Car{"Swift"}
	Status(obj)
}