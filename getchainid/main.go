package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/FactomProject/factom"
)

func main() {
	e := new(factom.Entry)
	os.Args = os.Args[1:]
	e.ExtIDs = make([][]byte, len(os.Args))
	for i, a := range os.Args {
		b, err := hex.DecodeString(a)
		if err != nil {
			panic(err)
		}
		e.ExtIDs[i] = b
	}

	c := factom.NewChain(e)
	fmt.Println(c.ChainID)
}
