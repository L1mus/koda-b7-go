package minitask1

import (
	"math"
)

func calculateArea(r float32) float32 {
	var area float32 = r * math.Phi
	return area
}

func calculateCircumference(r float32) float32 {
	var circumference float32 = 2 * math.Phi * r
	return circumference
}

func CalculateCircle(r float32) (float32, float32) {
	area := calculateArea(r)
	circumference := calculateCircumference(r)
	return area, circumference
}
