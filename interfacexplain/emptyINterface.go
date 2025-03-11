package interfacexplain

import "fmt"

func PrintValue(v interface{}){
	fmt.Println(v)
}

func EmptyInterface(){
	PrintValue(42)
	PrintValue("savana")
}


/**
* Here we have empty interface which is use for we can pass any data type value in argument of function
*/