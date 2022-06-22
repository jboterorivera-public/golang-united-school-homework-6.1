package golang_united_school_homework

import (
	"reflect"
	"testing"
)

func TestBoxCapacity(t *testing.T) {
	capacity := 7

	box := *NewBox(capacity)

	want := capacity
	got := box.shapesCapacity

	if got != want {
		t.Errorf("TestBoxCapacity got: %v, want: %v", got, want)
	}
}

func TestAddShapes(t *testing.T) {
	capacity := 2

	box := *NewBox(capacity)
	box.AddShape(Rectangle{Height: 4, Weight: 6})
	box.AddShape(Triangle{Side: 5})

	want := capacity
	got := len(box.shapes)

	if got != want {
		t.Errorf("TestAddShapes got: %v, want: %v", got, want)
	}
}

func TestAddShapeNoSpaceToAddShapes(t *testing.T) {
	capacity := 1

	box := *NewBox(capacity)
	box.AddShape(Rectangle{Height: 4, Weight: 6})

	want := errorNoSpaceForShapes
	got := box.AddShape(Triangle{Side: 5})

	if got.Error() != want.Error() {
		t.Errorf("TestAddShapeNoSpaceToAddShapes got: %v, want: %v", got, want)
	}
}

func TestGetByIndex(t *testing.T) {
	capacity := 3

	box := *NewBox(capacity)

	s := Rectangle{Height: 8, Weight: 10}

	box.AddShape(Rectangle{Height: 4, Weight: 6})
	box.AddShape(s)
	box.AddShape(Triangle{Side: 5})

	want := s
	got, _ := box.GetByIndex(1)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestGetByIndex got: %v, want: %v", got, want)
	}
}

func TestGetByIndexCountShapes(t *testing.T) {
	capacity := 3

	box := *NewBox(capacity)

	box.AddShape(Rectangle{Height: 4, Weight: 6})
	box.AddShape(Rectangle{Height: 8, Weight: 10})
	box.AddShape(Triangle{Side: 5})

	want := 3
	box.GetByIndex(1)
	got := len(box.shapes)

	if got != want {
		t.Errorf("TestGetByIndexCountShapes got: %v, want: %v", got, want)
	}
}

func TestGetByIndexOutOfIndex(t *testing.T) {
	capacity := 3

	box := *NewBox(capacity)

	box.AddShape(Rectangle{Height: 8, Weight: 10})
	box.AddShape(Triangle{Side: 5})

	want := errorElementOutOfIndex
	_, got := box.GetByIndex(2)

	if got.Error() != want.Error() {
		t.Errorf("TestGetByIndexOutOfIndex got: %v, want: %v", got, want)
	}
}

func TestExtractByIndexCompareShape(t *testing.T) {
	capacity := 3

	box := *NewBox(capacity)

	s := Rectangle{Height: 8, Weight: 10}

	box.AddShape(Rectangle{Height: 4, Weight: 6})
	box.AddShape(s)
	box.AddShape(Triangle{Side: 5})

	want := s
	got, _ := box.ExtractByIndex(1)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("TestExtractByIndexCompareShape got: %v, want: %v", got, want)
	}
}

func TestExtractByIndexCountShapes(t *testing.T) {
	capacity := 3

	box := *NewBox(capacity)

	box.AddShape(Rectangle{Height: 4, Weight: 6})
	box.AddShape(Rectangle{Height: 8, Weight: 10})
	box.AddShape(Triangle{Side: 5})

	want := 2
	box.ExtractByIndex(1)
	got := len(box.shapes)

	if got != want {
		t.Errorf("TestExtractByIndexCountShapes got: %v, want: %v", got, want)
	}
}

func TestExtractByIndexOutOfIndex(t *testing.T) {
	capacity := 3

	box := *NewBox(capacity)

	box.AddShape(Rectangle{Height: 8, Weight: 10})
	box.AddShape(Triangle{Side: 5})

	want := errorElementOutOfIndex
	_, got := box.ExtractByIndex(2)

	if got.Error() != want.Error() {
		t.Errorf("TestExtractByIndexOutOfIndex got: %v, want: %v", got, want)
	}
}

func TestReplaceByIndex(t *testing.T) {
	capacity := 3

	box := *NewBox(capacity)

	s := Rectangle{Height: 8, Weight: 10}
	sAux := Circle{Radius: 78}

	box.AddShape(Rectangle{Height: 8, Weight: 6})
	box.AddShape(s)
	box.AddShape(Triangle{Side: 5})

	want := s
	got, _ := box.ReplaceByIndex(1, sAux)

	if got != want {
		t.Errorf("TestReplaceByIndex got: %v, want: %v", got, want)
	}
}

func TestReplaceByIndexOutOfIndex(t *testing.T) {
	capacity := 3

	box := *NewBox(capacity)

	box.AddShape(Rectangle{Height: 8, Weight: 10})
	box.AddShape(Triangle{Side: 5})

	s := Circle{Radius: 78}

	want := errorElementOutOfIndex
	_, got := box.ReplaceByIndex(2, s)

	if got.Error() != want.Error() {
		t.Errorf("TestReplaceByIndexOutOfIndex got: %v, want: %v", got, want)
	}
}

func TestSumPerimeter(t *testing.T) {
	capacity := 3

	box := *NewBox(capacity)

	box.AddShape(Rectangle{Height: 8, Weight: 10})
	box.AddShape(Triangle{Side: 5})
	box.AddShape(Circle{Radius: 7})

	want := 94.9822971502571
	got := (float64)(0)
	for _, s := range box.shapes {
		got += s.CalcPerimeter()
	}

	if got != want {
		t.Errorf("TestSumPerimeter got: %v, want: %v", got, want)
	}
}

func TestSumArea(t *testing.T) {
	capacity := 3

	box := *NewBox(capacity)

	box.AddShape(Rectangle{Height: 8, Weight: 10})
	box.AddShape(Triangle{Side: 5})
	box.AddShape(Circle{Radius: 7})

	want := 244.76335757320533
	got := (float64)(0)
	for _, s := range box.shapes {
		got += s.CalcArea()
	}

	if got != want {
		t.Errorf("TestSumArea got: %v, want: %v", got, want)
	}
}

func TestRemoveCircles(t *testing.T) {
	capacity := 6

	b := *NewBox(capacity)

	b.AddShape(Rectangle{Height: 8, Weight: 10})
	b.AddShape(Triangle{Side: 5})
	b.AddShape(Circle{Radius: 7})
	b.AddShape(Rectangle{Height: 9, Weight: 11})
	b.AddShape(Triangle{Side: 6})
	b.AddShape(Circle{Radius: 8})

	b.RemoveAllCircles()

	for _, s := range b.shapes {
		_, ok := s.(Circle)
		if ok {
			t.Errorf("TestRemoveCircles, the box still contains Circles in its shapes")
		}
	}

	want := 4
	got := len(b.shapes)

	if got != want {
		t.Errorf("TestRemoveCircles got: %v, want: %v", got, want)
	}
}

func TestRemoveCirclesInShapesWithoutCircles(t *testing.T) {
	capacity := 4

	box := *NewBox(capacity)

	box.AddShape(Rectangle{Height: 8, Weight: 10})
	box.AddShape(Triangle{Side: 5})
	box.AddShape(Rectangle{Height: 9, Weight: 11})
	box.AddShape(Triangle{Side: 6})

	want := errorNoShapestoRemove
	got := box.RemoveAllCircles()

	if got.Error() != want.Error() {
		t.Errorf("TestRemoveCirclesInShapesWithoutCircles got: %v, want: %v", got, want)
	}
}
