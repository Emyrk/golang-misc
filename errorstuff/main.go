package main

import (
	"errors"
	"fmt"

	"golang.org/x/xerrors"
)

type CustomError struct {
	Value string
}

func (c CustomError) Error() string {
	return c.Value
}

func main() {
	var c CustomError

	err := cError()
	fmt.Println(errors.As(err, &c))
	fmt.Println(c)

	var c2 CustomError
	fmt.Println(xerrors.As(err, &c2))
	fmt.Println(c2)

}

func cError() error {
	return CustomError{
		Value: "Here is an error",
	}
}
