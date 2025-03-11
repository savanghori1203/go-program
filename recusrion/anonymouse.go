package recusrion

import "fmt"

func Anonymous(){
	var greet = func ()  {
		fmt.Println("Hello world!")
	}
	greet()
}