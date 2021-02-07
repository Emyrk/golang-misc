package main

import "fmt"

var a int = 0
var Exported int = 4

func printit() {
	var b int = 1
	fmt.Printf("a=%d b=%d, E=%d\n", a, b, Exported) // here is breakpoint

}

func main() {
	printit()
}
