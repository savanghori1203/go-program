package encodinganddecoding

import (
	"encoding/json"
	"fmt"
)

type Formate struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Price    int    `json:"price"`
}

func Unmarshal() {
	jsonString := `{"name" : "savan", "platform": "salesmate", "price" : 2345}`

	jsonData := []byte(jsonString)

	var course Formate

	err := json.Unmarshal(jsonData, &course)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(course)
}
