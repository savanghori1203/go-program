package errorexplain

import "fmt"

type error interface {
	Error() string
}

type MyError struct {
	Code    int
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func Something() error {
	return MyError{Code : 404 , Message: "Not found"}
}

func CustomErrorTesting(){
	err := Something()
	fmt.Println(err)
}