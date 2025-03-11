package structexplain

import "fmt"

type Anonymous struct {
	personalDetails struct {
		Name string
		Age  int
	}
	GPA float64
}

func Three() {
  an := Anonymous{
	personalDetails: struct{
       Name string
	   Age int
	}{
		Name : "savan",
		Age : 4,
	},
	GPA: 9.8,
  }

  fmt.Printf("Anonymouse struct detail : %+v\n",an)
}
