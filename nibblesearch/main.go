package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"github.com/FactomProject/factomd/database/databaseOverlay"
	"github.com/FactomProject/factomd/database/hybridDB"
)

const level string = "level"
const bolt string = "bolt"

func main() {
	var (
		loc = flag.String("db", "", "DB Location")
		db  = flag.String("t", "level", "DBType, level or bolt")
	)

	flag.Parse()

	fmt.Println("Usage:")
	fmt.Println("anaylzedb -t level/bolt -db DBFileLocation")
	fmt.Println("Program will analyzeblocks")

	levelBolt := *db

	if levelBolt != level && levelBolt != bolt {
		fmt.Println("\nFirst argument should be `level` or `bolt`")
		os.Exit(1)
	}
	path := *loc

	var dbase *hybridDB.HybridDB
	var err error
	if levelBolt == bolt {
		dbase = hybridDB.NewBoltMapHybridDB(nil, path)
	} else {
		dbase, err = hybridDB.NewLevelMapHybridDB(path, false)
		if err != nil {
			panic(err)
		}
	}

	dbo := databaseOverlay.NewOverlay(dbase)
	var _ = dbo

	head, err := dbo.FetchDBlockHead()
	if err != nil {
		fmt.Println(err)
		return
	}
	top := head.GetDatabaseHeight()

	_ = os.Remove("nibbles.")
	nibbles, err := os.OpenFile("nibbles.csv", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(nibbles)
	_ = w.Write([]string{"Height", "Since Last Nibble"})

	last := uint32(0)
	for i := uint32(0); i < top; i++ {
		dblock, err := dbo.FetchDBlockByHeight(i)
		if err != nil {
			panic(err)
		}

		keymr := dblock.GetKeyMR()
		if keymr.Bytes()[0]&0xF0 == 0 {
			ht := dblock.GetDatabaseHeight()
			_ = w.Write([]string{fmt.Sprintf("%d", ht), fmt.Sprintf("%d", ht-last)})
			fmt.Println(keymr.String(), ht, ht-last)
			last = dblock.GetDatabaseHeight()
		}
	}

}
