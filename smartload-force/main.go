package main

import (
	//"encoding/hex"
	"bytes"
	"flag"
	"fmt"
	"strings"
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
	totalTrans *int32

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
	s.totalTrans = new(int32)

	return s
}

func (sc *StatCollector) PrintStatCollector() {
	tk := time.NewTicker(Each)
	start := time.Now()
	for _ = range tk.C {
		tc := atomic.LoadInt32(sc.totalChains)
		te := atomic.LoadInt32(sc.totalEntries)
		vc := atomic.LoadInt32(sc.verifyChains)
		ve := atomic.LoadInt32(sc.verifyEntries)
		tt := atomic.LoadInt32(sc.totalTrans)
		fmt.Printf("%f TPS |  C: %d/%d  --  E: %d/%d\n", float64(tt)/time.Since(start).Seconds(), vc, tc, ve, te)
	}
}

func (sc *StatCollector) NewTransaction(v int32) {
	atomic.AddInt32(sc.totalTrans, v)
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

func cycleServer(hosts []string) {
	factom.SetFactomdServer(hosts[0])
	i := 0
	max := len(hosts)
	ticker := time.NewTicker(time.Second * 1)
	for _ = range ticker.C {
		if i >= max {
			i = 0
		}
		factom.SetFactomdServer(hosts[i])
		i++
	}
}

func main() {
	var (
		hosts       = flag.String("hosts", "localhost:8088", "Factomd location")
		blktime     = flag.Int("blktime", 10, "Block time in second")
		sleepAmount = flag.Int("s", 5, "Sleep duration")
	)

	flag.Parse()
	fmt.Printf("Hosts:%s\nBlktime:%d\nSleep:%d\n", *hosts, *blktime, *sleepAmount)

	BlockTime = time.Duration(*blktime) * time.Second

	hostArr := strings.Split(*hosts, " ")
	if len(hostArr) == 1 {
		factom.SetFactomdServer(hostArr[0])
	} else {
		go cycleServer(hostArr)
	}
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
		go makeRandomChain(sc)
		time.Sleep(time.Duration(*sleepAmount) * time.Second)
	}
}

func makeRandomChain(sc *StatCollector) {
	hash := random.RandByteSliceOfLen(5)
	c, err := makeChain()
	if err != nil {
		fmt.Println(err)
		return
	}
	sc.NewTransaction(1)

	amt := random.RandIntBetween(0, 25)
	ents, err := addEntries(c, amt)
	if err != nil {
		fmt.Println(err)
		return
	}
	sc.NewTransaction(int32(amt))

	sle := (BlockTime / 10)
	for i := 0; i < 12; i++ {
		var _ = hash
		// fmt.Printf("[%x] %d/%d -- %fs left\n", hash, i, 12, float64(12-i)*sle.Seconds())
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
	_, err := factom.CommitChain(c, SandEC)
	if err != nil {
		return nil, err
	}

	_, err = factom.RevealChain(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func addEntries(c *factom.Chain, entrycount int) ([]*EntryHolder, error) {
	var ents []*EntryHolder
	for i := 0; i < entrycount; i++ {
		e := randomEntry()
		e.ChainID = c.ChainID
		_, err := factom.CommitEntry(e, SandEC)
		if err != nil {
			return nil, err
		}

		entryhash, err := factom.RevealEntry(e)
		if err != nil {
			return nil, err
		}
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
