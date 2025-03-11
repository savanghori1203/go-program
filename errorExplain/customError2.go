package errorexplain

import "fmt"

type CustErr struct {
	Message string
}

func (c CustErr) Error() string {
	return "Number can't be divided by zero"
}

func DivideTest(a int, b int) (int, error) {
	if b == 0 {
		return 0, CustErr{}
	}
	return a / b, nil
}

func TestingDivide() {
	value, err := DivideTest(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}


/**
How does it work?
 In Go, any type that implements the "Error() string" method satisfies the error interface.
 When you return CustErr{} in DivideTest(), it is treated as an error type.
 If an error is returned and printed with fmt.Println(err), Go automatically calls the Error() method. 
*/

/**
Here automatically convert into error interface 
*/