package main

import (
	"flag"
	"fmt"

	"github.com/Emyrk/factom-raw"
)

func main() {
	var (
		path = flag.String("p", "/root/.factom/m2/custom-database/ldb/CUSTOM/factoid_level.db", "Path to db")
		ht   = flag.Int("ht", 0, "Height to grab")
	)

	flag.Parse()

	db := factom_raw.NewDBReader("level", *path)
	dblock, err := db.FetchDBlockByHeight(uint32(*ht))
	if err != nil {
		panic(err)
	}

	fmt.Println(dblock)
	data, err := dblock.MarshalBinary()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\v", data)

}
