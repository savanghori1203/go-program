package channelexplain

import "fmt"

func ProducerMessage(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func ConsumerMessage(ch chan int) {
	for item := range ch {
		fmt.Println(item)
	}
}

func MessageParisngGoroutine() {
	ch := make(chan int)
	go ProducerMessage(ch)
	ConsumerMessage(ch)
}
