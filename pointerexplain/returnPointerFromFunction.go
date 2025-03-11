package pointerexplain

import "fmt"

func ReturnPointerFromFunction(){
   ptr := Test()
   fmt.Println(ptr)
}

func Test() *string {
   name := "savan"

   return &name
}