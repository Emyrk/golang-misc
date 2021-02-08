package pgp

import (
	"crypto"
	"crypto/rand"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
)

func main() {

}

func gen() {
	config := packet.Config{
		Rand:        rand.Reader,
		DefaultHash: crypto.SHA512,
		//DefaultCipher:          crypto.,
		Time:                   nil,
		DefaultCompressionAlgo: 0,
		CompressionConfig:      nil,
		S2KCount:               0,
		RSABits:                0,
	}
	var _ = config
	e, _ := openpgp.NewEntity("name", "", "a@gmail.com", nil)

	//e.Serialize()


}
