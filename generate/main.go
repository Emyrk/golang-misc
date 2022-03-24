package main

import (
	"fmt"

	"golang.org/x/tools/go/packages"
)

func main() {
	cfg := &packages.Config{
		Mode: packages.NeedTypes,
	}

	fmt.Println(cfg)
}
