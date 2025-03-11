package advancetoipcexplain

import "fmt"

func DivisionTest(num1, num2 int) {

	// if num2 is 0, program is terminated due to panic
	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		result := num1 / num2
		fmt.Println("Result: ", result)
	}
}

func PanicOut() {

	DivisionTest(4, 2)
	DivisionTest(8, 0)
	DivisionTest(2, 8)

}

/**
We use the panic statement to immediately end the execution of the program. 
If our program reaches a point where it cannot be recovered due to some major errors, it's best to use panic.
*/