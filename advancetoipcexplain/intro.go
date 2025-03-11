package advancetoipcexplain

import "fmt"

/**
We use defer to delay the execution of functions that might cause an error.
The panic statement terminates the program immediately and recover is used to recover the message during panic.

We use the defer statement to prevent the execution of a function until all other functions execute. 
*/

func ExpDefer() {
	defer fmt.Println("tree")
	fmt.Println("one")
	fmt.Println("two")
}
