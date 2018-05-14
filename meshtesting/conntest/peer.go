package main

import (
	"log"

	"github.com/weaveworks/mesh"
)

// Peer encapsulates state and implements mesh.Gossiper.
// It should be passed to mesh.Router.NewGossip,
// and the resulting Gossip registered in turn,
// before calling mesh.Router.Start.
type SimplePeer struct {
	send   mesh.Gossip
	quit   chan struct{}
	logger *log.Logger
}

// Return a copy of our complete state.
func (p *SimplePeer) Gossip() (complete mesh.GossipData) {
	p.logger.Printf("Gossip => complete %v", complete)
	return complete
}

// Merge the gossiped data represented by buf into our state.
// Return the state information that was modified.
func (p *SimplePeer) OnGossip(buf []byte) (delta mesh.GossipData, err error) {
	return NewPacket(buf), nil
}

// Merge the gossiped data represented by buf into our state.
// Return the state information that was modified.
func (p *SimplePeer) OnGossipBroadcast(src mesh.PeerName, buf []byte) (received mesh.GossipData, err error) {
	return NewPacket(buf), nil
}

// Merge the gossiped data represented by buf into our state.
func (p *SimplePeer) OnGossipUnicast(src mesh.PeerName, buf []byte) error {
	packet := NewPacket(buf)
	p.logger.Printf("OnGossipUnicast %s => complete %v", src, string(packet.Encode()[0]))
	return nil
}
