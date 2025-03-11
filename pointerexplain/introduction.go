package pointerexplain

import "fmt"

func Introduction1(){
  var x int = 10

  var ptr *int = &x
  fmt.Println(ptr)
}