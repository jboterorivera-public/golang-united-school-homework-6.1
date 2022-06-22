package golang_united_school_homework

import "fmt"

type CustomError struct {
	Message string
	Err     error
}

func (c CustomError) Error() string {
	return c.Message
}

func (c CustomError) Unwrap() error {
	return c.Err
}

func customError(e error) (c CustomError) {
	c = CustomError{Message: e.Error()}
	c.Err = fmt.Errorf("Error %w", e)

	return c
}
