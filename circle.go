package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (receiver Circle) CalcPerimeter() float64 {
	return (2 * math.Pi) * receiver.Radius
}

func (receiver Circle) CalcArea() float64 {
	var diameter = math.Pow(receiver.Radius, 2)
	var result = diameter * math.Pi

	return result
}
