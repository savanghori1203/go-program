package interfacexplain

import "fmt"

//value, ok := interfaceValue.(TargetType)
/**
  if same tye value will be return adn ok will be true
  if fail of will be false we have to hanlde it for both
*/

func typeAssertion(v interface{}){
	str , ok := v.(string)
	if ok {
		fmt.Println("string type", str)
	}else{
		fmt.Println("Assertion failed")
	}
}

func TypeAssertionInInterface(){
	typeAssertion(42)
	typeAssertion("savana")
}
