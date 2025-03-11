package mapexplain

import "fmt"

func Nested(){
	nested := map[string]map[string]string{
		"savan" : {
			"surname" : "Ghori",
			"work" : "developer",
			"status" : "true", 
		},
		"tirth" : {
			"surname" : "Shah",
			"work" : "developer",
			"status" : "false", 
		},
	}

	fmt.Println(nested)

	 _,value := nested["savan"]

	if value {
      fmt.Println("key exist",value)
	}else{
		fmt.Println("key don't exist")
	}

	fmt.Println(nested["savan"])
}