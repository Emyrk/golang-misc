package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/FactomProject/factomd/common/messages"
)

func main() {
	a := new(messages.DirectoryBlockSignature)
	v, err := hex.DecodeString(os.Args[1])
	if err != nil {
		panic(err)
	}

	nd, err := a.UnmarshalBinaryData(v)
	if err != nil {
		panic(err)
	}

	if len(nd) > 0 {
		fmt.Printf("Length of bytes remaining is %d", len(nd))
	}

	fmt.Println(a.String())
	str, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(str))

}
