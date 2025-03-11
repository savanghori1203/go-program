package goroutineexplain

import (
	"fmt"
	"time"
)

func DisplayMessage(message string) {
	fmt.Println(message)
}

func Goroutine() {
	go DisplayMessage("process 1")

	DisplayMessage("process 2")
	time.Sleep(time.Second * 3)
}

/**
Both functions together asynchronously.
We can convert a regular function to a goroutine by calling the function using the go keyword.
We get message print still the main goroutine is executing if it will end we cant get (we have to put time wait to run concurrent other task)
Main goroutine is default and run alway
*/

/**
if we have multiple goroutine we need communication we use channel to comunicate independly
*/
