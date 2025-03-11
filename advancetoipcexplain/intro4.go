package advancetoipcexplain

import "fmt"

// recover function to handle panic
func handlePanic() {

	// detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println("RECOVER", a)
	}

}

func DivisionDefer(num1, num2 int) {

	// execute the handlePanic even after panic occurs
	defer handlePanic()

	// if num2 is 0, program is terminated due to panic
	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		result := num1 / num2
		fmt.Println("Result: ", result)
	}
}

func UseRecoveryAndPanic() {

	DivisionDefer(4, 2)
	DivisionDefer(8, 0)
	DivisionDefer(8, 2)

}


/**
Once panic occurs in DivisionDefer(8, 0), recover() will handle it, 
but execution won't continue to the next function call (DivisionDefer(2, 8)) because the panic only stops the current goroutine, not the entire program.
*/

/**
last DivisionDefer(2,8) run beacuase it will stop only current goroutine
Since panic only affects the goroutine where it occurs, 
calling UseRecoveryAndPanic() inside main() ensures that the program runs fully.
*/