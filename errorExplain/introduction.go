package errorexplain

import (
	"errors"
	"fmt"
)

func Divide(a,b float32) (float32, error) {
   if b == 0 {
	 return 0, errors.New("can't be divide by zero")
   }
   return a/b, nil
}

func ErrorHandling(){
	result,err := Divide(10,0)
	if err != nil {
		fmt.Println("Error",err)
		return
	}
	fmt.Println("Resut is", result)
}

