package main

import (
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/FactomProject/factomd/common/messages/msgsupport"

	"github.com/FactomProject/factomd/common/messages"

	"github.com/FactomProject/factom"
	"github.com/FactomProject/factomd/common/entryCreditBlock"
)

const ECPriv = "Es2XT3jSxi1xqrDvS5JERM3W3jh1awRHuyoahn3hbQLyfEi1jvbq"

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 400; i++ {
		Run()
	}
}

func Run() {
	extids := make([][]byte, rand.Intn(5))
	for i := range extids {
		extids[i] = make([]byte, rand.Intn(10))
		rand.Read(extids[i])
	}

	e := new(factom.Entry)
	e.ExtIDs = extids
	e.Content = []byte("Burning")
	c := factom.NewChain(e)

	addr, _ := factom.GetECAddress(ECPriv)
	req, _ := ComposeChainCommit(c, addr)
	res, err := factom.SendFactomdRequest(req)
	fmt.Println(res, err)
}

// ComposeChainCommit creates a JSON2Request to commit a new Chain via the
// factomd web api. The request includes the marshaled MessageRequest with the
// Entry Credit Signature.
func ComposeChainCommit(c *factom.Chain, ec *factom.ECAddress) (*factom.JSON2Request, error) {
	buf := new(bytes.Buffer)

	// 1 byte version
	buf.Write([]byte{0})

	// 6 byte milliTimestamp
	buf.Write(milliTime())

	e := c.FirstEntry

	// 32 byte ChainID Hash
	if p, err := hex.DecodeString(c.ChainID); err != nil {
		return nil, err
	} else {
		// double sha256 hash of ChainID
		buf.Write(shad(p))
	}

	// 32 byte Weld; sha256(sha256(EntryHash + ChainID))
	if cid, err := hex.DecodeString(c.ChainID); err != nil {
		return nil, err
	} else {
		s := append(e.Hash(), cid...)
		buf.Write(shad(s))
	}

	// 32 byte Entry Hash of the First Entry
	buf.Write(e.Hash())

	// 1 byte number of Entry Credits to pay
	if d, err := factom.EntryCost(e); err != nil {
		return nil, err
	} else {
		d = 20
		buf.WriteByte(byte(d))
	}

	// 32 byte Entry Credit Address Public Key + 64 byte Signature
	sig := ec.Sign(buf.Bytes())
	buf.Write(ec.PubBytes())
	buf.Write(sig[:])

	data := buf.Bytes()

	commit := entryCreditBlock.NewCommitChain()
	err := commit.UnmarshalBinary(data)
	if err != nil {
		panic(err)
	}

	msg := new(messages.CommitChainMsg)
	msg.CommitChain = commit

	data, err = msg.MarshalBinary()

	messages.General = new(msgsupport.GeneralFactory)
	_, _ = messages.General.UnmarshalMessage(data)

	params := SendRawMessageRequest{Message: hex.EncodeToString(data)}
	req := factom.NewJSON2Request("send-raw-message", factom.APICounter(), params)

	fmt.Println(commit.GetEntryHash().String())
	return req, nil
}

type messageRequest struct {
	Message string `json:"message"`
}

type SendRawMessageRequest struct {
	Message string `json:"message"`
}

// milliTime returns a 6 byte slice representing the unix time in milliseconds
func milliTime() (r []byte) {
	buf := new(bytes.Buffer)
	t := time.Now().UnixNano()
	m := t / 1e6
	binary.Write(buf, binary.BigEndian, m)
	return buf.Bytes()[2:]
}

// shad Double Sha256 Hash; sha256(sha256(data))
func shad(data []byte) []byte {
	h1 := sha256.Sum256(data)
	h2 := sha256.Sum256(h1[:])
	return h2[:]
}

// sha52 Sha512+Sha256 Hash; sha256(sha512(data)+data)
func sha52(data []byte) []byte {
	h1 := sha512.Sum512(data)
	h2 := sha256.Sum256(append(h1[:], data...))
	return h2[:]
}
