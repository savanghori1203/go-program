package arrayexplain

import "fmt"

var Arraylist = [4]int{1,2,3,4}
var arr = [3]string{0:"savan", 2:"hiren"}
func One(){
	fmt.Println(Arraylist)
	fmt.Println(len(Arraylist))
	fmt.Println(arr)
	fmt.Printf("Array detail : %+v\n",arr)
}