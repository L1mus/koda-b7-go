package main

import (
	"fmt"
)

func main() {
	insertNumber()
}

func insertNumber() {
	var newSlice []int = make([]int, 0, 7)
	newElement := 88
	numbers := []int{50, 75, 66, 20, 32, 90}

	newSlice = append(newSlice, numbers[0:3]...)
	newSlice = append(newSlice, newElement)
	newSlice = append(newSlice, numbers[3:6]...)

	//using func inserts
	//numbers = slices.Insert(numbers, 3, 88)
	//
	//for i, v := range numbers {
	//	fmt.Printf("angka ke-%d : %d\n", i, v)
	//}

	for i, v := range newSlice {
		fmt.Printf("angka ke-%d : %d\n", i, v)
	}

}
