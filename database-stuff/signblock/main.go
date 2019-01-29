package main

import (
	"flag"
	"io/ioutil"
	"os"

	"encoding/json"

	"fmt"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/messages"
	"github.com/FactomProject/factomd/common/primitives"
)

type Keys struct {
	Keys []string `json:"keys"`
}

func main() {
	var (
		fileName = flag.String("f", "", "Dbstate filename")
		keys     = flag.String("k", "", "File with keys to sign")
		out      = flag.String("o", "", "Output file")
	)

	flag.Parse()

	dbstateFile, err := os.OpenFile(*fileName, os.O_RDONLY, 0777)
	panicE(err)
	defer dbstateFile.Close()
	data, err := ioutil.ReadAll(dbstateFile)
	panicE(err)

	dbstate := new(messages.DBStateMsg)
	_, err = dbstate.UnmarshalBinaryData(data)
	panicE(err)

	keysFile, err := os.OpenFile(*keys, os.O_RDONLY, 0777)
	panicE(err)
	defer keysFile.Close()
	data, err = ioutil.ReadAll(keysFile)
	panicE(err)
	k := Keys{}
	err = json.Unmarshal(data, &k)
	panicE(err)

	var signingKeys []*primitives.PrivateKey
	for _, s := range k.Keys {
		priv, err := primitives.NewPrivateKeyFromHex(s)
		panicE(err)
		signingKeys = append(signingKeys, priv)
	}

	data, err = dbstate.DirectoryBlock.GetHeader().MarshalBinary()
	panicE(err)

	var sigList []interfaces.IFullSignature
	for _, sec := range signingKeys {
		sigList = append(sigList, sec.Sign(data))
	}

	dbstate.SignatureList.List = sigList
	dbstate.SignatureList.Length = uint32(len(sigList))

	dbstate.MarshalBinary()
	// data, err := m.DirectoryBlock.GetHeader().MarshalBinary()
	//dbstate.SignatureList =

	fmt.Println(dbstate)
	file, err := os.OpenFile(*out, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err = dbstate.MarshalBinary()
	if err != nil {
		panic(err)
	}
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}

}

func panicE(err error) {
	if err != nil {
		panic(err)
	}
}
