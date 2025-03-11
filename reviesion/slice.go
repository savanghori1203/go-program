package reviesion

import "fmt"

func SliceLearning() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(slice))

	newSlice := slice[:1]
	fmt.Println(newSlice)

	slicenew := ReturnSlice(slice)
	fmt.Println(slicenew)

	for _,value := range slice {

		if value == 4{
			continue
		}
		fmt.Println(value)
	}

	// this is using the loop
	for i :=0 ; i < len(slice); i++{
		if slice[i] == 4 {
			continue
		}
		fmt.Println(slice[i])
	}
}


func ReturnSlice(lice []int) []int {
	return lice[:2]
}
