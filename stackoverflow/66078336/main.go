package main

import (
	"fmt"
)

type runner interface {
	isRunner() bool
	run()
}

type person struct {
}

type program struct {
}

func (person) isRunner() bool { return true }
func (p person) run()         {}
func (p program) run()        {}

func main() {
	var per interface{} = person{}
	var prog interface{} = program{}

	_, ok := per.(runner)
	fmt.Printf("Person is a runner: %t\n", ok)

	_, ok = prog.(runner)
	fmt.Printf("Program is a runner: %t\n", ok)
}
