package channelexplain

import (
	"fmt"
	"time"
)

func ReceivedMessage3(name chan string) {
	fmt.Println(<-name)
	fmt.Println("Message Received! Received Operation Complete")
}

func TetingChannel3() {
	name := make(chan string)

	go ReceivedMessage3(name)

	fmt.Println("Befor send the data to the name channel")

	name <- "sedn data to channel"

	fmt.Println("After send the data to the name channel")


	time.Sleep(5 * time.Second)
}

/**
Block conditon : 2
 we can't send the message though channel from main goroutine it will block execution of ReceivedMessage3
 if we comment name <- "sedn data to channel"
*/