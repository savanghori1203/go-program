package stringexplain

import (
	"fmt"
)

func Looping() {
	message := "Hello i am savan ghori"

	for _, values := range message {
		fmt.Printf( "%c\n" , values)
	}
}
