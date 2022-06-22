package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (receiver Rectangle) CalcPerimeter() float64 {
	return 2 * (receiver.Height + receiver.Weight)
}

func (receiver Rectangle) CalcArea() float64 {
	return receiver.Height * receiver.Weight
}
