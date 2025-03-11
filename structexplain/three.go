package structexplain

import "fmt"

type Engineer struct {
	Name string
	Age  int
	Company Company
}

type Company struct {
	Name string
	Location string
	Experience int
}

func (e Company) Details() string {
	return "company environment is good for you"
}

func (e Engineer) Hello() string {
   return "hello"
}

func TestCompany() {
	details := Engineer{
		Name: "savan",
		Age: 30,
		Company: Company{
			Name: "Google",
			Location: "clf",
			Experience: 5,
		},
	}

	fmt.Printf("Company Details: %v\n", details)
}