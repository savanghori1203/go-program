package package1

import (
	"fmt"
)

func PublicFunc(){
	fmt.Println("public function")
}

func privateFunc(){
	fmt.Println("private function")
}

func AccessPrivateFunc(){
	PublicFunc()
}