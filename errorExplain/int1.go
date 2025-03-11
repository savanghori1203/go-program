package errorexplain

import (
	"errors"
	"fmt"
)

func ErrorExplain() (string, error) {
	message := "test"

	myError := errors.New("wrong Message")

	if message != "svana" {
		return "not found", myError
	}
	return message, nil
}

func ReturnCustomError(){
	message,err := ErrorExplain()
	fmt.Println(message)
	fmt.Println(err)
}

/**
In Go, we can use the New() function to handle an error. This function is defined inside the errors package and allows us to create our own error message.
*/

