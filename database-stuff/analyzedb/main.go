package main

import (
	"fmt"
	"os"

	"flag"

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
	fmt.Println("top", head.GetDatabaseHeight())

	// 206024 is missing fblock
	ht := uint32(206024)

	fblock, _ := dbo.FetchFBlockByHeight(ht)
	fmt.Println("f", fblock.GetDatabaseHeight(), fblock.GetKeyMR().String())
	fmt.Println(fblock.String())

	dblock, _ := dbo.FetchDBlockByHeight(ht)
	fmt.Println("d", dblock.GetDatabaseHeight(), dblock.GetKeyMR().String())

	fmt.Println("\tf->", dblock.GetDBEntries()[2].GetKeyMR().String())

	//fmt.Println(dbo.Fetc(206572))

}
