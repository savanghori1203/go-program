package mapexplain

import "fmt"

func Introduction2(){
	newMap := make(map[string]string)

	newMap["name"] = "savan"
	newMap["job"] = "developer" 
	
	fmt.Println(newMap)
}

/**
* we can't do like this for initializing the slice like

NOTE:
   var newMap map[string]strng
   it will throw error 

   if we want to do this way we need to do initialize the map through make() function

   newMap = make(map[string]string)

   methode 2: 

   newMap := map[string]string {
   value
   }
*/