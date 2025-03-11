package channelexplain

import (
	"fmt"
	"time"
)

func ReceivedMessage2(name chan string) {
	time.Sleep(1 * time.Second)
	name <- "message added"
	fmt.Println("Message sent! Send Operation Complete")
}

func TetingChannel2() {
	name := make(chan string)

	go ReceivedMessage2(name)

	fmt.Println(<-name)

	fmt.Println("block still we don't received the data from goroutine")

	time.Sleep(5 * time.Second)
}


/**
Block condition :1 
block the go routine execution still the we don't received the channel message 
if we <-name in TetingChannel2 function comment then dircet it will print last fmt message 
*/