package main

import (
	"fmt"
	log "github.com/FactomProject/logrus"
)

func main() {
	fmt.Println("Hello, playground")
	var identLogger = log.WithFields(log.Fields{"subpack": "identity"})
	identLogger.Error("Hi")
	another := log.WithFields(log.Fields{"Hello": "s"})
	// identLogger.Data
	identLogger = identLogger.WithFields(another.Data)
	log.WithFields(identLogger.Data).Error("ASD")
	log.WithFields(log.Fields{
		"package":  "messages",
		"func":     "Validate",
		"message":  "DBState",
		"dbheight": 5,
		"llheight": 5,
		"vm":       0,
		"signer":   "88AAEE",
		"prevmr":   "ABCDEF",
		"hash":     "123BCD",
	}).Error("Failed to validate")

	fmt.Printf("%v\n", identLogger)
}
