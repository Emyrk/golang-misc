package main

import (
	"fmt"

	"github.com/FactomProject/factomd/common/primitives"
)

func main() {
	p := primitives.RandomPrivateKey()
	fmt.Printf("Priv: %s\n", p.PrivateKeyString())
	fmt.Printf("Pub : %s\n", p.PublicKeyString())

}
