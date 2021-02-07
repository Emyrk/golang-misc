package main

import (
	"foo"
	"fmt"
	"os"
)

func main() {
	foo.Bar()
	Foo()
}

var bad_name int

func Foo() error {
	fmt.Println(bad_name)
	return fmt.Errorf("ASD")
}

func PrintAllDirectories(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		fmt.Println(info.Name())
	}
	return nil
}
