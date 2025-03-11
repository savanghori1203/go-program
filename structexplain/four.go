package structexplain

import "fmt"

type Number struct {
	X int
	Y int
}

func (p Number) Addition() int {
	ans := p.X + p.Y
	return ans
}

func (p Number) Multilpication() int {
	ans := p.X * p.Y
	return ans
} 

func Output(){
	answer := Number{ X : 1, Y: 4}

	add := answer.Addition()
	mut := answer.Multilpication()

	fmt.Printf("Addition of %d and %d is : %d\n", answer.X , answer.Y , add)
	fmt.Printf("Multiplication of %d and %d is : %d\n", answer.X , answer.Y , mut)

}