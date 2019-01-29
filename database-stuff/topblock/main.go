package main

import (
	"fmt"
	"os"

	"flag"

	"github.com/FactomProject/factomd/database/databaseOverlay"
	"github.com/FactomProject/factomd/database/hybridDB"
	"github.com/FactomProject/factomd/state"
)

const level string = "level"
const bolt string = "bolt"

func main() {
	var (
		loc = flag.String("db", "", "DB Location")
		db  = flag.String("t", "level", "DBType, level or bolt")
		ht  = flag.Int("h", 0, "Height to be grabbed")
	)

	flag.Parse()

	fmt.Println("Usage:")
	fmt.Println("FixBlockHeads -t level/bolt -db DBFileLocation")
	fmt.Println("Program will reset the block heads to the highest valid DBlock")

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
	s := new(state.State)
	s.DB = dbo

	height := uint32(*ht)
	if *ht == 0 {
		head, _ := dbo.FetchDirectoryBlockHead()
		height = head.GetDatabaseHeight()
	}
	fmt.Println("Grabbing", height)

	s.ProcessLists = new(state.ProcessLists)
	s.ProcessLists.DBHeightBase = height
	p1 := new(state.ProcessList)
	p1.DBHeight = height + 1
	p1.DBSignatures = []state.DBSig{}
	s.ProcessLists.Lists = []*state.ProcessList{p1}

	dbstate, err := s.LoadDBState(height)
	if err != nil {
		panic(err)
	}
	fmt.Println(dbstate)
	file, err := os.OpenFile(fmt.Sprintf("dbstate_%d.dbstate", height), os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := dbstate.MarshalBinary()
	if err != nil {
		panic(err)
	}
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}
