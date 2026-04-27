package main

import (
	"fmt"
)

func main() {
	createTriangel(5)
}

func createTriangel(n int) {
	for i := range n {
		for j := 0; j <= n-i; j++ {
			fmt.Printf(" ")
		}
		for k := 0; k <= 2*i; k++ {
			fmt.Printf("*")

		}
		fmt.Printf("\n")

	}
}
