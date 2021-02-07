package main

import (
	"errors"
	"fmt"
)

type MyErrorMessage struct {
	Line     int
	Position int
	Msg      string
}

func (err MyErrorMessage) Error() string {
	return fmt.Sprintf("ERR: line %d::%d %s", err.Line, err.Position, err.Msg)
}

func foo() error {
	return MyErrorMessage{Line: 1, Position: 2, Msg: "help!"}
}

func main() {
	err := fmt.Errorf("An error: %w", foo())

	if myErr, ok := err.(MyErrorMessage); ok {
		fmt.Println("Type Cast", myErr.Line, myErr.Position)
	}

	var mem MyErrorMessage
	if errors.As(err, &mem) {
		fmt.Println("Errors.As", mem.Line, mem.Position)
	}
}
