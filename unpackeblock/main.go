package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/FactomProject/factomd/common/entryBlock"
)

func main() {
	flag.Parse()
	m := flag.Args()[0]
	b, err := hex.DecodeString(m)
	if err != nil {
		panic(err)
	}

	dbs := new(entryBlock.EBlock)
	err = dbs.UnmarshalBinary(b)
	if err != nil {
		panic(err)
	}

	jb, err := dbs.JSONByte()
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer

	err = json.Indent(&buf, jb, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(buf.Bytes()))

}
