package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (receiver Triangle) CalcPerimeter() float64 {
	return receiver.Side * float64(SidesTriangle)
}

func (receiver Triangle) CalcArea() float64 {
	var calcSides = math.Pow(receiver.Side, 2)
	var triangleConst = math.Sqrt(float64(SidesTriangle)) / 4
	var result = (calcSides * triangleConst)

	return result
}
