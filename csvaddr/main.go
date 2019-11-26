package main

import (
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Factom-Asset-Tokens/factom"
)

func main() {
	out, err := os.OpenFile("out.csv", os.O_CREATE|os.O_WRONLY, 0777)
	p(err)
	defer out.Close()

	in, err := os.OpenFile("in.csv", os.O_RDONLY, 0777)
	p(err)
	defer in.Close()

	outR := csv.NewWriter(out)
	inR := csv.NewReader(in)
	record, err := inR.Read()
	for len(record) > 0 {
		p(err)
		d, _ := hex.DecodeString(record[0])
		var addr factom.FAAddress
		copy(addr[:], d[:])
		record[0] = addr.String()

		amt, err := strconv.ParseInt(record[1], 10, 64)
		p(err)

		record = append(record, FactoshiToFactoid(amt))

		err = outR.Write(record)
		p(err)

		record, err = inR.Read()
	}

}

func p(err error) {
	if err != nil {
		panic(err)
	}
}

// FactoshiToFactoid converts a uint64 factoshi ammount into a fixed point
// number represented as a string
func FactoshiToFactoid(i int64) string {
	d := i / 1e8
	r := i % 1e8
	ds := fmt.Sprintf("%d", d)
	rs := fmt.Sprintf("%08d", r)
	rs = strings.TrimRight(rs, "0")
	if len(rs) > 0 {
		ds = ds + "."
	}
	return fmt.Sprintf("%s%s", ds, rs)
}
