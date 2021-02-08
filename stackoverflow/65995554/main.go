package main

import (
	"fmt"
	"reflect"
)

var _ SomeInterface = SomeModel{}

type SomeModel struct {
	SomeField string
}

func (sm SomeModel) GetSomeField() string {
	return sm.SomeField
}

type SomeInterface interface {
	GetSomeField() string
}

func NewSomeInterface() SomeInterface {
	return &SomeModel{}
}

func methodP(someInterface *SomeInterface) {
	someInterface.GetSomeField() // no access to method
}

func method(someInterface SomeInterface) {
	someInterface.GetSomeField() // access allowed
}

func main() {
	s := NewSomeInterface()
	fmt.Println(reflect.TypeOf(s))
}
