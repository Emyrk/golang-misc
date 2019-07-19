package benchmarks

import (
	"context"
	"sync"
)

// BalanceMap enables CRUD of balances in a threadsafe manner.
type BalanceMap struct {
	balances map[[32]byte]int64

	queries chan *BalanceQuery

	chanPool chan chan int64

	quit chan struct{}

	sync.RWMutex
}

type BalanceQuery struct {
	qt      int
	Address [32]byte

	Balance int64
	done    chan struct{}
	*sync.WaitGroup

	ret chan int64
}

func NewBalanceQuery() *BalanceQuery {
	q := new(BalanceQuery)

	return q
}

func NewBalanceMap() *BalanceMap {
	b := new(BalanceMap)
	b.quit = make(chan struct{})
	b.queries = make(chan *BalanceQuery, 10)
	b.balances = make(map[[32]byte]int64)
	b.chanPool = make(chan chan int64, 100)
	for i := 0; i < cap(b.chanPool); i++ {
		b.chanPool <- make(chan int64, 1)
	}

	return b
}

func (bm *BalanceMap) Close() {
	close(bm.quit)
Loop:
	for {
		select {
		case q := <-bm.queries:
			bm.Respond(q)
		default:
			break Loop
		}
	}
	close(bm.queries)
}

func (bm *BalanceMap) Closed() bool {
	select {
	case _, open := <-bm.quit:
		return !open
	default:
		return false
	}
}

func (bm *BalanceMap) Respond(q *BalanceQuery) {
	switch q.qt {
	case 0:
		q.Balance = bm.balances[q.Address]
		//close(q.done)
		q.WaitGroup.Done()
	case 1:
		fallthrough
	case 2:
		bal := bm.balances[q.Address]
		q.ret <- bal
	}
}

// Serve will enable the serving
func (bm *BalanceMap) Serve() {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	for {
		select {
		case <-bm.quit:
		case q := <-bm.queries:
			bm.Respond(q)
		}
	}
}

func (bm *BalanceMap) GetBalance(addr [32]byte, qt int) int64 {
	switch qt {
	case -1:
		return bm.GetDirectBalance(addr)
	case 0:
		return bm.GetBalanceSharedQuery(addr)
	case 1:
		return bm.GetBalanceChanBackQuery(addr)
	case 2:
		return bm.GetBalanceSharedChanBackQuery(addr)
	}
	return -1
}

// GetBalanceChanBackQuery shares the query struct
func (bm *BalanceMap) GetBalanceSharedChanBackQuery(addr [32]byte) int64 {
	q := NewBalanceQuery()
	q.qt = 2
	q.Address = addr
	q.ret = <-bm.chanPool // Grab the chan to return

	bm.queries <- q // Send the query
	bal := <-q.ret  // Wait for the response
	bm.chanPool <- q.ret

	return bal
}

// GetBalanceChanBackQuery shares the query struct
func (bm *BalanceMap) GetBalanceChanBackQuery(addr [32]byte) int64 {
	q := NewBalanceQuery()
	q.qt = 1
	q.Address = addr
	q.ret = make(chan int64, 1)
	bm.queries <- q // Send the query
	return <-q.ret  // Wait for the response
}

// GetBalanceSharedQuery shares the query struct
func (bm *BalanceMap) GetBalanceSharedQuery(addr [32]byte) int64 {
	q := NewBalanceQuery()
	//q.done = make(chan struct{})
	q.qt = 0
	q.Address = addr
	q.WaitGroup = new(sync.WaitGroup)
	q.WaitGroup.Add(1)

	bm.queries <- q // Send the query
	q.WaitGroup.Wait()
	return q.Balance // Return the result
}

// _GetBalance should only be used by tests
func (bm *BalanceMap) GetDirectBalance(addr [32]byte) int64 {
	bm.Lock()
	defer bm.Unlock()
	return bm.balances[addr]
}

//func (bm *BalanceMap) _GetBalance(addr [32]byte) <-chan int64 {
//	c := make(chan int64, 1)
//	c <- bm.balances[addr]
//	close(c)
//	return c
//}

func (bm *BalanceMap) SetDirectBalance(addr [32]byte, amt int64) {
	bm.balances[addr] = amt
}

func r() {
	ctx := context.Background()
	ctx.Done()
}
