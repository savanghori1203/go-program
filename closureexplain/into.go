package closureexplain

import "fmt"

func Closure() func() string{
   name := "savan"
   return func() string{
	   return ("Hello" + " " + name)
   }
   
}

func ReturnValue(){
	result := Closure()
	fmt.Println(result())
}


/**
Here anonymouse function check for it parent score for name variable for findinf this it called closure
*/