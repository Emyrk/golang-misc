package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"time"

	"github.com/FactomProject/factom"
	"github.com/FactomProject/factomd/common/directoryBlock"
	"github.com/FactomProject/factomd/common/entryBlock"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"github.com/FactomProject/factomd/database/databaseOverlay"
	"github.com/FactomProject/factomd/database/hybridDB"
)

var CheckFloating bool
var UsingAPI bool
var FixIt bool

const level string = "level"
const bolt string = "bolt"

var blockamount = 10
var blocktime = 600 // In seconds

func main() {
	var (
		// Use API by default
		apiHost = flag.String("h", "localhost:8088", "Change the host API")
	)
	flag.Parse()

	fmt.Println("Usage: countentries -h localhost:8088")
	fmt.Println("Program will count total chains and entries")

	apireader := NewAPIReader(*apiHost)
	count(apireader)

	// fmt.Printf("Program finished.\n  Total Chains:  %d\n  Total Entries: %d\n", chains, entries)
}

func count(reader Fetcher) {
	top, err := reader.FetchDBlockHead()
	if err != nil {
		panic(err)
	}
	height := top.GetDatabaseHeight()

	bottom := height - uint32(blockamount)
	if height <= uint32(blockamount) {
		bottom = 0
	}

	fmt.Printf("Calculating rate over blocks %d to %d.\n", bottom, height)
	entryCount := 0
	entryBlockCount := 0

	for i := int(bottom); i <= int(height); i++ {
		perblockEntries := 0
		block, err := reader.FetchDBlockByHeight(uint32(i))
		if err != nil {
			fmt.Println(err)
			time.Sleep(100 * time.Millisecond)
			i--
			continue
		}

		eblocks := block.GetEBlockDBEntries()
		entryBlockCount += len(eblocks)
		// for _, ebh := range eblocks {
		for c := 0; c < len(eblocks); c++ {
			ebh := eblocks[c]
			eb, err := reader.FetchEBlock(ebh.GetKeyMR())
			if err != nil {
				fmt.Println(err)
				time.Sleep(100 * time.Millisecond)
				c--
				continue
			}

			perblockEntries += len(eb.GetEntryHashes())
			entryCount += len(eb.GetEntryHashes())
		}
		fmt.Printf(" Block %d: %d Entries at %f/s\n", i, perblockEntries, float64(perblockEntries)/float64(blocktime))
	}

	fmt.Printf("Totals -- Entries: %d\n", entryCount)
	fmt.Printf("Per Block Average -- Entries: %f\n", float64(entryCount)/float64(blockamount))
	fmt.Printf("PerSecond -- Entries: %f\n", float64(entryCount)/float64(blockamount*blocktime))
}

type Fetcher interface {
	FetchDBlockHead() (interfaces.IDirectoryBlock, error)
	FetchDBlockByHeight(dBlockHeight uint32) (interfaces.IDirectoryBlock, error)
	//FetchDBlock(hash interfaces.IHash) (interfaces.IDirectoryBlock, error)
	FetchHeadIndexByChainID(chainID interfaces.IHash) (interfaces.IHash, error)
	FetchEBlock(hash interfaces.IHash) (interfaces.IEntryBlock, error)
	SetChainHeads(primaryIndexes, chainIDs []interfaces.IHash) error
}

func NewDBReader(levelBolt string, path string) *databaseOverlay.Overlay {
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
	return dbo
}

type APIReader struct {
	location string
}

func NewAPIReader(loc string) *APIReader {
	a := new(APIReader)
	a.location = loc
	factom.SetFactomdServer(loc)

	return a
}

func (a *APIReader) SetChainHeads(primaryIndexes, chainIDs []interfaces.IHash) error {
	return nil
}

func (a *APIReader) FetchEBlock(hash interfaces.IHash) (interfaces.IEntryBlock, error) {
	raw, err := factom.GetRaw(hash.String())
	if err != nil {
		return nil, err
	}
	return rawBytesToEblock(raw)
}

func (a *APIReader) FetchDBlockHead() (interfaces.IDirectoryBlock, error) {
	head, err := factom.GetDBlockHead()
	if err != nil {
		return nil, err
	}
	raw, err := factom.GetRaw(head)
	if err != nil {
		return nil, err
	}
	return rawBytesToDblock(raw)
}

func (a *APIReader) FetchDBlockByHeight(dBlockHeight uint32) (interfaces.IDirectoryBlock, error) {
	raw, err := factom.GetBlockByHeightRaw("d", int64(dBlockHeight))
	if err != nil {
		return nil, err
	}

	return rawRespToBlock(raw.RawData)
}

func (a *APIReader) FetchHeadIndexByChainID(chainID interfaces.IHash) (interfaces.IHash, error) {
	resp, err := factom.GetChainHead(chainID.String())
	if err != nil {
		return nil, err
	}
	return primitives.HexToHash(resp)
}

func rawBytesToDblock(raw []byte) (interfaces.IDirectoryBlock, error) {
	dblock := directoryBlock.NewDirectoryBlock(nil)
	err := dblock.UnmarshalBinary(raw)
	if err != nil {
		return nil, err
	}
	return dblock, nil
}

func rawBytesToEblock(raw []byte) (interfaces.IEntryBlock, error) {
	eblock := entryBlock.NewEBlock()
	err := eblock.UnmarshalBinary(raw)
	if err != nil {
		return nil, err
	}
	return eblock, nil
}

func rawRespToBlock(raw string) (interfaces.IDirectoryBlock, error) {
	by, err := hex.DecodeString(raw)
	if err != nil {
		return nil, err
	}
	return rawBytesToDblock(by)
}
