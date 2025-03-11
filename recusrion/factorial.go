package recusrion

import "fmt"

func Factorial(num int) int {
	if num == 0 {
		return 1
	}else{
		return num * Factorial(num-1)
	}
}

func AllFactorial(){
	var name int
	fmt.Print("Please inter value\n")
	fmt.Scanf("%d",&name)
	result := Factorial(name)
	fmt.Println(result)
}