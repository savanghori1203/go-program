package reviesion

import "fmt"

func StringLerning() string{
    return "heloo world"
}

func StringCallingOutside(){
	result := StringLerning()
	fmt.Println(result)
}

