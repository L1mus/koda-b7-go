package main

import (
	"fmt"
	"math"
)

func main() {
	luas, keliling := calculateCircle(2.5)
	fmt.Printf("Keliling Lingkaran %.2f , Luas Lingkaran %.2f\n", keliling, luas)

	area := calculateArea(2.5)
	circumference := calculateCircumference(2.5)
	fmt.Printf("Keliling Lingkaran %.2f , Luas Lingkaran %.2f", circumference, area)

}

func calculateCircle(r float32) (float32, float32) {
	var area float32 = r * math.Phi
	var circumference float32 = 2 * math.Phi * r
	return area, circumference
}

func calculateArea(r float32) float32 {
	var area float32 = r * math.Phi
	return area
}

func calculateCircumference(r float32) float32 {
	var circumference float32 = 2 * math.Phi * r
	return circumference
}
