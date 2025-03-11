package encodinganddecoding

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Price    int    `json:"price"`
}

func MarshalData() {
	course := Course{
		Name:     "savam",
		Platform: "experro",
		Price:    23456,
	}

	fmt.Println(course)

	jsonData, err := json.Marshal(course)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(string(jsonData))

}

/**
🔹 What Happens in Marshal?
The Go struct is converted into a JSON string.
Struct field names are mapped to JSON keys.
json.Marshal() returns a byte slice ([]byte).
We use string(jsonData) to print it as a string.

*/
