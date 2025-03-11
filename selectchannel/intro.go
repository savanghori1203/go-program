package selectchannel

import (
	"fmt"
	"time"
)

func ProcessMessage(ch chan string) {
	time.Sleep(time.Second * 3)
	ch <- "savan"

}

func ProcessNumber(ch chan int) {
	// time.Sleep(time.Second * 3)
	ch <- 17
}

func SelectChannel() {
	num := make(chan int)
	message := make(chan string)

	go ProcessNumber(num)
	go ProcessMessage(message)

	time.Sleep(time.Second * 3)

	select {

	case firstChannel := <-num:
		fmt.Println("Channle data", firstChannel)

	case secondChannel := <-message:
		fmt.Println("Channel data", secondChannel)

	default:
		fmt.Println("Wait!! Channels are not ready for execution")
	}

}

/**
Here we get random value but mimimu atleast one received still if i have not add wait in main goroutine adn that why we use select
We set default value if we don't received any value
*/
