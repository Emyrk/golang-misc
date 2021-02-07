//nolint
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("expected an argument to specify which directory to search")
	}

	root := flag.Args()[0]
	fmt.Println("Searching", root)

	filepath.Walk(root, PrintAllDirectories)
}

func PrintAllDirectories(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		fmt.Println(info.Name())
	}
	return nil
}
