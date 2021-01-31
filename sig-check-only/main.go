package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/FactomProject/factomd/common/primitives"
)

func main() {
	var (
		pub   = flag.String("pub", "", "Public key")
		sig   = flag.String("sig", "", "Signature")
		msg = flag.String("msg", "", "Message")
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
	sigH, err := hex.DecodeString(*sig)
	panicerr(err, "1")
	pubH, err := hex.DecodeString(*pub)
	panicerr(err, "2")

	err = s.SetSignature(sigH)
	panicerr(err, "3")
	s.SetPub(pubH)

	msgData, err := hex.DecodeString(*msg)
	panicerr(err, "2")
	fmt.Println(s.Verify(msgData))
	fmt.Println(s.Verify(msgData))


}

func panicerr(err error, where string) {
	if err != nil {
		fmt.Println(where)
		panic(err)
	}
}
