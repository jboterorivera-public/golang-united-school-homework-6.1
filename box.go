package golang_united_school_homework

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	totalShapes := len(b.shapes)

	if totalShapes+1 > b.shapesCapacity {
		return customError(errorNoSpaceForShapes)
	}

	b.shapes = append(b.shapes, shape)

	return nil
}

func (b *box) getShape(index int) (Shape, error) {
	if index > (b.shapesCapacity-1) || index >= len(b.shapes) {
		return nil, customError(errorElementOutOfIndex)
	}

	shape := b.shapes[index]

	if shape == nil {
		return nil, customError(errorIndexWithoutShape)
	}

	return shape, nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	shape, err := b.getShape(i)

	if err != nil {
		return nil, err
	}

	return shape, nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	shape, err := b.getShape(i)

	if err != nil {
		return nil, err
	}

	boxAux := *NewBox(b.shapesCapacity)
	for i2, s := range b.shapes {
		if i != i2 {
			boxAux.shapes = append(boxAux.shapes, s)
		}
	}

	*b = boxAux

	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	oldShape, err := b.getShape(i)

	if err != nil {
		return nil, err
	}

	b.shapes[i] = shape

	return oldShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	acum := (float64)(0)

	for i := range b.shapes {
		shape, err := b.getShape(i)

		if err != nil {
			continue
		}

		acum += shape.CalcPerimeter()
	}

	return acum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	acum := (float64)(0)

	for i := range b.shapes {
		shape, err := b.getShape(i)

		if err != nil {
			continue
		}

		acum += shape.CalcArea()
	}

	return acum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	boxAux := *NewBox(b.shapesCapacity)

	counter := 0
	for _, s := range b.shapes {
		_, ok := s.(Circle)
		if !ok {
			boxAux.shapes = append(boxAux.shapes, s)
		} else {
			counter++
		}
	}

	if counter == 0 {
		return customError(errorNoShapestoRemove)
	}

	*b = boxAux

	return nil
}
