package channelexplain

import "fmt"

func ReceivedMessage4(ch chan int){
 for i := 1 ;i < 5 ;i++ {
	ch <- i
 }
 close(ch) // we can't close it will wait
}

func LoopigInChannel(){
	ch := make(chan int)

	go ReceivedMessage4(ch)

	for val := range ch{
		fmt.Println(val)
	}
	//range allow as to take the value one by one
}