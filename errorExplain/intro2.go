package errorexplain

import "fmt"

func CreateManuallyError() {
	age := -14

	error := fmt.Errorf("%d is negative\nAge can't be negative", age)

	if age < 0 {
		fmt.Println(error)
	} else {
		fmt.Printf("Age %d", age)
	}
}

/**
We can also handle Go errors using the Errorf() function. Unlike, New(), we can format the error message using Errorf().

This function is present inside the fmt package, so we can directly use this if we have imported the fmt package.
*/
