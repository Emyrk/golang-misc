package main

import (
	//"encoding/hex"
	"bytes"
	"flag"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/FactomProject/factom"
	"github.com/FactomProject/factom/wallet"
	"github.com/FactomProject/factomd/common/primitives/random"
)

var _ = fmt.Println

// FA address with factoids
var (
	// ECSandStr is EC1zPkB9FHJEpDgGLACTGECipjLKzYHM1pd9ywCcRcBF2GrZP2AT
	ECSandStr = "Es3AreQSueoZyg6XrsWQG7NAQNw6RWZytBquPZmJbPK86P8C4mK7"

	// FASandStr is FA2jK2HcLnRdS94dEcU27rF3meoJfpUcZPSinpb7AwQvPRY6RL1Q
	FASandStr = "Fs3E9gV6DXsYzf7Fqx1fVBQPQXV695eP3k5XbmHEZVRLkMdD9qCK"
	SandFA    *factom.FactoidAddress

	SandEC *factom.ECAddress

	Rate      uint64
	BlockTime time.Duration
	Each      = time.Second * 5
)

func init() {
	//f, _ := hex.DecodeString(FASandStr)
	//e, _ := hex.DecodeString(ECSandStr)

	SandFA, _ = factom.GetFactoidAddress(FASandStr)
	SandEC, _ = factom.GetECAddress(ECSandStr)
}

type EntryHolder struct {
	Entry     *factom.Entry
	EntryHash string
}

type StatCollector struct {
	totalChains  *int32
	totalEntries *int32

	verifyChains  *int32
	verifyEntries *int32
}

func NewStatCollector() *StatCollector {
	s := new(StatCollector)
	s.totalChains = new(int32)
	s.totalEntries = new(int32)
	s.verifyChains = new(int32)
	s.verifyEntries = new(int32)

	return s
}

func (sc *StatCollector) PrintStatCollector() {
	tk := time.NewTicker(Each)
	for _ = range tk.C {
		tc := atomic.LoadInt32(sc.totalChains)
		te := atomic.LoadInt32(sc.totalEntries)
		vc := atomic.LoadInt32(sc.verifyChains)
		ve := atomic.LoadInt32(sc.verifyEntries)
		fmt.Printf("C: %d/%d  --  E: %d/%d\n", vc, tc, ve, te)
	}
}

func (sc *StatCollector) VerifyChain(v int32) {
	atomic.AddInt32(sc.verifyChains, v)
}

func (sc *StatCollector) VerifyEntry(v int32) {
	atomic.AddInt32(sc.verifyEntries, v)
}

func (sc *StatCollector) AddChain(v int32) {
	atomic.AddInt32(sc.totalChains, v)
}

func (sc *StatCollector) AddEntry(v int32) {
	atomic.AddInt32(sc.totalEntries, v)
}

func main() {
	var (
		host    = flag.String("s", "localhost:8088", "Factomd location")
		blktime = flag.Int("blktime", 10, "Block time in second")
	)

	flag.Parse()
	BlockTime = time.Duration(*blktime) * time.Second

	factom.SetFactomdServer(*host)
	rate, err := factom.GetRate()
	panicerr(err)
	Rate = rate

	wal, _ := wallet.NewMapDBWallet()

	panicerr(wal.InsertFCTAddress(SandFA))
	panicerr(wal.InsertECAddress(SandEC))

	fundEC(wal)
	sc := NewStatCollector()
	go sc.PrintStatCollector()

	for {
		//go makeRandomChain(sc)
		go func() {
			_, err := makeChain()
			sc.AddChain(1)
			if err == nil {
				fmt.Println(err)
				sc.VerifyChain(1)
			}
		}()

		_, err := makeChain()
		sc.AddChain(1)
		if err == nil {
			fmt.Println(err)
			sc.VerifyChain(1)
		}
		// time.Sleep(10 * time.Millisecond)
	}
}

func makeRandomChain(sc *StatCollector) {
	hash := random.RandByteSliceOfLen(5)
	c, err := makeChain()
	if err != nil {
		fmt.Println(err)
		return
	}

	amt := 1000 //random.RandIntBetween(250, 500)
	ents, err := addEntries(c, amt)
	if err != nil {
		fmt.Println(err)
		return
	}

	sle := (BlockTime / 10)
	for i := 0; i < 12; i++ {
		fmt.Printf("[%x] %d/%d -- %fs left\n", hash, i, 12, float64(12-i)*sle.Seconds())
		time.Sleep(sle)
	}
	sc.AddChain(1)
	sc.AddEntry(int32(amt))

	errs, fc, failed := verify(c, ents, sc)
	for _, e := range errs {
		fmt.Println(e)
	}

	time.Sleep(BlockTime + (BlockTime / 10))
	errs, _, _ = verify(fc, failed, sc)
	for _, e := range errs {
		fmt.Println(e)
	}
}

func verify(c *factom.Chain, ents []*EntryHolder, sc *StatCollector) ([]error, *factom.Chain, []*EntryHolder) {
	var failed []*EntryHolder
	var fc *factom.Chain
	var errs []error
	if c != nil {
		fe, err := factom.GetFirstEntry(c.ChainID)
		if err != nil {
			errToAdd := fmt.Errorf("Err with chain %s: %s", c.ChainID, err.Error())
			fc = c
			errs = append(errs, errToAdd)
		}

		if err := compareEntries(fe, c.FirstEntry); err != nil {
			errs = append(errs, err)
		} else {
			sc.VerifyChain(1)
		}
	}

	for _, e := range ents {
		entry, err := factom.GetEntry(e.EntryHash)
		if err != nil {
			errToAdd := fmt.Errorf("Err with entry %s: %s", e.EntryHash, err.Error())
			failed = append(failed, e)
			errs = append(errs, errToAdd)
			continue
		}

		if err := compareEntries(entry, e.Entry); err != nil {
			errs = append(errs, err)
			continue
		}
		sc.VerifyEntry(1)
	}
	return errs, fc, failed
}

func compareEntries(a *factom.Entry, b *factom.Entry) error {
	if a == nil || b == nil {
		return fmt.Errorf("One is nil")
	}
	if bytes.Compare(a.Content, b.Content) != 0 {
		return fmt.Errorf("Content does not match. A is len %d, B is len %d\n  -- (a) Chain is: %s, EntryHash is %x\n  -- (b) Chain is: %s, EntryHash is %x",
			len(a.Content), len(b.Content), a.ChainID, a.Hash(), b.ChainID, b.Hash())
	}

	return nil
}

func makeChain() (*factom.Chain, error) {
	c := factom.NewChain(randomEntry())
	txid, err := factom.CommitChain(c, SandEC)
	if err != nil {
		return nil, fmt.Errorf("Error CommitChain: %s", err.Error())
	}

	err, stat, stat2 := waitonAck(txid, true)
	if err != nil {
		return nil, err
	}

	rtxid, err := factom.RevealChain(c)
	if err != nil {
		return nil, fmt.Errorf("Error RevealChain: %s", err.Error())
	}

	fmt.Println("factom-cli get allentries", c.ChainID)

	err, stat, stat2 = waitonAck(rtxid, false)
	if err != nil {
		return nil, err
	}
	var _, _ = stat, stat2

	_, err = factom.GetChainHead(c.ChainID)
	if err != nil {
		fmt.Println("CH: ", stat, stat2, err.Error(), rtxid, c.ChainID)
		return c, fmt.Errorf("ASd")
	}

	return c, nil
}

/*
	if !*fflag {
		if _, err := waitOnCommitAck(txid); err != nil {
			errorln(err)
			return
		}
	}
	// reveal entry
	hash, err := factom.RevealEntry(e)
	if err != nil {
		errorln(err)
		return
	}
	if !*fflag {
		if _, err := waitOnRevealAck(txid); err != nil {
			errorln(err)
			return
		}
	}
*/

func waitonAck(txid string, com bool) (error, string, string) {
	for {
		if com {
			s, err := factom.EntryCommitACK(txid, "")
			if err != nil {
				return fmt.Errorf("Error WaitOnAck: %s", err.Error()), "", ""
			}

			if (s.CommitData.Status != "Unknown") && (s.CommitData.Status != "NotConfirmed") {
				return nil, s.CommitData.Status, s.EntryData.Status
			}
		} else {
			s, err := factom.EntryACK(txid, "")
			if err != nil {
				return fmt.Errorf("Error WaitOnAck: %s", err.Error()), "", ""
			}

			if (s.EntryData.Status != "Unknown") && (s.EntryData.Status != "NotConfirmed") {
				return nil, s.CommitData.Status, s.EntryData.Status
			}
		}
		// time.Sleep(time.Millisecond / 100)
	}
}

func addEntries(c *factom.Chain, entrycount int) ([]*EntryHolder, error) {
	var ents []*EntryHolder
	for i := 0; i < entrycount; i++ {
		e := randomEntry()
		e.ChainID = c.ChainID
		txid, err := factom.CommitEntry(e, SandEC)
		if err != nil {
			return nil, fmt.Errorf("Error CommitEntry: %s", err.Error())
		}

		err, stat, stat2 := waitonAck(txid, true)
		if err != nil {
			return nil, err
		}

		entryhash, err := factom.RevealEntry(e)
		if err != nil {
			return nil, fmt.Errorf("Error RevealEntry: %s", err.Error())
		}

		err, stat, stat2 = waitonAck(entryhash, false)
		if err != nil {
			return nil, err
		}
		var _, _ = stat, stat2

		eh := new(EntryHolder)
		eh.Entry = e
		eh.EntryHash = entryhash
		ents = append(ents, eh)
	}
	return ents, nil
}

func randomEntry() *factom.Entry {
	e := new(factom.Entry)
	l := random.RandIntBetween(0, 8000)
	e.Content = random.RandByteSliceOfLen(l)
	e.ExtIDs = make([][]byte, random.RandIntBetween(2, 7))
	for i := range e.ExtIDs {
		e.ExtIDs[i] = random.RandByteSliceOfLen(random.RandIntBetween(0, 100))
	}

	return e
}

func fundEC(wal *wallet.Wallet) {
	panicerr(wal.NewTransaction("buy"))
	amt := uint64(20 * 1e8)
	panicerr(wal.AddECOutput("buy", "EC1zPkB9FHJEpDgGLACTGECipjLKzYHM1pd9ywCcRcBF2GrZP2AT", amt))
	panicerr(wal.AddInput("buy", "FA2jK2HcLnRdS94dEcU27rF3meoJfpUcZPSinpb7AwQvPRY6RL1Q", amt))
	panicerr(wal.AddFee("buy", "FA2jK2HcLnRdS94dEcU27rF3meoJfpUcZPSinpb7AwQvPRY6RL1Q", Rate))
	panicerr(wal.SignTransaction("buy", false))
	tx, err := wal.ComposeTransaction("buy")
	panicerr(err)

	_, err = factom.SendFactomdRequest(tx)
	panicerr(err)
}

func panicerr(err error) {
	if err != nil {
		panic(err)
	}
}
