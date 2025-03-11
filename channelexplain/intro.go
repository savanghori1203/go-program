package channelexplain

import "fmt"

func ReceivedMessage(num chan int, message chan string) {
   num <- 15
   message <- "savan testing"
}

func TetingChannel() {
	number := make(chan int)
	message := make(chan string)

	go ReceivedMessage(number, message)

	fmt.Println(<-number)
	fmt.Println(<-message)
}

/**
<- use to send data in channel 
Once we create a channel, we can send and receive data between different goroutines through the channel.
*/