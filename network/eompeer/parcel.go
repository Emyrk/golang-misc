package main

import (
	"github.com/FactomProject/factomd/p2p"
	"time"
)

type P2PParcel struct {
	Parcel *p2p.Parcel
	sent   time.Time
}

func NewP2Parcel(p *p2p.Parcel) *P2PParcel {
	pp := new(P2PParcel)
	pp.Parcel = p
	return pp
}