package advancetoipcexplain

import "fmt"

func MultipalDefer(){
	defer fmt.Println("one")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
}

/**
When we use multiple defer in a program, the order of execution of the defer statements will be LIFO (Last In First Out).
*/