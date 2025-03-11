package interfacexplain

import "fmt"


//Interface

//data type
 /**
 * if we are calling the the api and in that we don't know the result of API
 * some time it give the id and some time id and name so we have to define the api result variable which store the data type to Interface
 * After define the type Interface we can assign the type of result by doing mashal adn unmarshal and looping over using rage loop
 * we can assing type of that particaular result what we get 
 */ 


//method
/**
* code will be clean and readable form
* in belove code struct define the the methode of it adn we made on interface which define every say must be return the string as result
*/


//withot interfaces
// type Cat struct {}

// func (c Cat) Say() string {
// 	return "meow"
// }

// type Dog struct {}

// func (d Dog) Say() string {
//  return "woof"
// }

// func Introduction(){
// 	c := Cat{}
// 	fmt.Println(c.Say())
// 	d := Dog{}
// 	fmt.Println(d.Say())
// }


//with interfaces

type Voice interface{
	Say() string
}

type Cat struct {}

func (c Cat) Say() string {
	if 1>0{
		return "voice form the cat"
	}
	return "meow"
}

type Dog struct {}

func (d Dog) Say() string {
if 4>5{
	return "voice from the dog"
}
 return "woof"
}

func Introduction(){
	c := Cat{}
	d := Dog{}
	voice := []Voice{c,d}

	fmt.Println(voice)
	for _,v := range voice{
		fmt.Println(v.Say())
	}
	
}


//advantage using interfaces

//polymorphism	
//data abtraction
//redable code
//understandable code
//testing make easy



