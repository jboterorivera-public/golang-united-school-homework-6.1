package golang_united_school_homework

import "errors"

var (
	errorNoSpaceForShapes  = errors.New("No space to store more shapes")
	errorElementOutOfIndex = errors.New("Element out of index")
	errorIndexWithoutShape = errors.New("The specified index doesn't have a stored shape")
	errorNoShapestoRemove  = errors.New("There are no shapes to remove with the specified type")
)
