package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/FactomProject/factom"
	"github.com/FactomProject/factomd/common/adminBlock"
	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/directoryBlock"
	"github.com/FactomProject/factomd/common/primitives"
)

func main() {
	var (
		pub   = flag.String("pub", "", "Public key")
		sig   = flag.String("sig", "", "Signature")
		serv  = flag.String("s", "localhost:8088", "Factomd location")
		keymr = flag.String("k", "", "KeyMr")
		next  = flag.String("n", "", "Next keymr of admin block")
	)

	flag.Parse()

	//h := uint32(*height)
	if *pub == "" || *sig == "" {
		if *next == "" {
			fmt.Println("Must specifiy a sig and pub")
			return
		}
	}

	s := new(primitives.Signature)
	if *next == "" {
		sigH, err := hex.DecodeString(*sig)
		panicerr(err, "1")
		pubH, err := hex.DecodeString(*pub)
		panicerr(err, "2")

		err = s.SetSignature(sigH)
		panicerr(err, "3")
		s.SetPub(pubH)
	}

	factom.SetFactomdServer(*serv)
	r, err := factom.GetRaw(*keymr)
	panicerr(err, "4")

	d := new(directoryBlock.DirectoryBlock)
	err = d.UnmarshalBinary(r)
	panicerr(err, "5")

	msg, err := d.Header.MarshalBinary()
	panicerr(err, "6")

	if *next != "" {
		ar, err := factom.GetRaw(*next)
		panicerr(err, "7")
		a := new(adminBlock.AdminBlock)
		err = a.UnmarshalBinary(ar)
		panicerr(err, "8")

		ents := a.GetABEntries()
		for _, e := range ents {
			if e.Type() == constants.TYPE_DB_SIGNATURE {
				dbsig := e.(*adminBlock.DBSignatureEntry)
				fmt.Println(dbsig.PrevDBSig.Verify(msg))
			}
		}
	} else {
		fmt.Println(s.Verify(msg))
	}
}

func panicerr(err error, where string) {
	if err != nil {
		fmt.Println(where)
		panic(err)
	}
}
