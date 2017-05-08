package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/FactomProject/factom"
	"github.com/FactomProject/factomd/common/adminBlock"
	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/directoryBlock"
	// "github.com/FactomProject/factomd/common/primitives"
)

func main() {
	var (
		serv   = flag.String("s", "localhost:8088", "Factomd location")
		height = flag.Int("h", 0, "Check this height")
		end    = flag.Int("e", 0, "End height")
	)

	flag.Parse()
	e := *end
	if e == 0 {
		e = *height + 1
	}

	factom.SetFactomdServer(*serv)

	for i := *height; i < e; i++ {
		verify(uint32(i))
	}

}

func verify(height uint32) {
	d := getDblockByHeight(height)
	a := getAblockByHeight(height + 1)
	msg, err := d.GetHeader().MarshalBinary()
	panicerr(err, "DblockHeader marshal")

	good := 0
	total := 0
	pre := ""
	for _, e := range a.GetABEntries() {
		if e.Type() == constants.TYPE_DB_SIGNATURE {
			dbsig := e.(*adminBlock.DBSignatureEntry)
			if !dbsig.PrevDBSig.Verify(msg) {
				pre := "!!"
				//fmt.Printf("--Bad Sig --\nID: %s\nSig: %x\n", dbsig.IdentityAdminChainID.String(), dbsig.PrevDBSig.GetSigBytes())
			} else {
				good++
			}
			total++
		}
	}
	fmt.Printf("%s >> %d/%d Good for %d.\n", pre, good, total, height)
}

func getDblockByHeight(height uint32) *directoryBlock.DirectoryBlock {
	dheight, err := factom.GetDBlockByHeight(int64(height))
	panicerr(err, "Dblock Fetch")
	d := new(directoryBlock.DirectoryBlock)
	r, err := hex.DecodeString(dheight.RawData)
	panicerr(err, "[D] HexDecode")

	err = d.UnmarshalBinary(r)
	panicerr(err, "UnmarshalDBlock")
	return d
}

func getAblockByHeight(height uint32) *adminBlock.AdminBlock {
	dheight, err := factom.GetABlockByHeight(int64(height))
	panicerr(err, "Ablock Fetch")
	a := new(adminBlock.AdminBlock)
	r, err := hex.DecodeString(dheight.RawData)
	panicerr(err, "[A] HexDecode")

	err = a.UnmarshalBinary(r)
	panicerr(err, "UnmarshalABlock")
	return a
}

func getABlock(keymr string) *adminBlock.AdminBlock {
	r, err := factom.GetRaw(keymr)
	panicerr(err, "GetRawAdmin")
	a := new(adminBlock.AdminBlock)
	err = a.UnmarshalBinary(r)
	panicerr(err, "UnmarshalABlock")
	return a
}

func panicerr(err error, where string) {
	if err != nil {
		fmt.Println(where)
		panic(err)
	}
}
