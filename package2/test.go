package package2

import (
    "fmt"
    "gotest/package1"
)

func AccessPublicFunc(){
    fmt.Println("Accessing public function from package2")
    package1.PublicFunc()
}



