package goroutineexplain

import (
	"fmt"
	"time"
)

func DisplayMessage2(message string) {
	fmt.Println(message)
}

func Goroutine2() {
	go DisplayMessage("process 1")
	go DisplayMessage2("process 2")
	go DisplayMessage2("process 3")

	time.Sleep(time.Second * 3)
}


/**
run concurrentyl and output is unexpected in form of index , here we wait for some time till all the execution in completed
main routine complete the  we can't do 
using go routine we can achive concurrentyl without using the os level thread
*/