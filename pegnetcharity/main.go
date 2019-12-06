package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Factom-Asset-Tokens/factom"
	"github.com/pegnet/pegnetd/fat/fat2"
	"github.com/pegnet/pegnetd/srv"
)

var inputAddr = "Fs2U6Mgxp6SYnPKRQzohhq7joiNdT6e9v7hJxcJLnVWL3g9pkiKC" // "FA2a544m3jShujvrVfyjN3mBZNq5oxSxT2XozSxGHf8uw9uR5xvT"

var TransactionChain [32]byte = [32]byte{0xcf, 0xfc, 0xe0, 0xf4, 0x09, 0xeb,
	0xba, 0x4e, 0xd2, 0x36, 0xd4, 0x9d, 0x89, 0xc7, 0x0e, 0x4b, 0xd1, 0xf1,
	0x36, 0x7d, 0x86, 0x40, 0x2a, 0x33, 0x63, 0x36, 0x66, 0x83, 0x26, 0x5a,
	0x24, 0x2d}

func main() {
	var err error
	cl := srv.NewClient()
	list := globalRich(cl, 1000)

	//for i := len(list) - 1; i >= 0; i-- {
	//	fmt.Println(list[i].Address)
	//}

	var batch fat2.TransactionBatch
	batch.Version = 1

	es, _ := factom.NewEsAddress("Es2ULrhz84Nb6J8c8L566i3h37D36LZtEg1bmN59gvGHVN6ngRGY")
	c := factom.Bytes32(TransactionChain)
	batch.Entry.ChainID = &c
	fs, _ := factom.NewFsAddress(inputAddr)

	var tx fat2.Transaction
	tx.Input.Amount = uint64(len(list))
	tx.Input.Address = fs.FAAddress()
	tx.Input.Type = fat2.PTickerPEG
	tx.Transfers = make([]fat2.AddressAmountTuple, len(list))

	var totalOut uint64 = 0
	for i := len(list) - 1; i >= 0; i-- {
		tx.Transfers[i].Amount = uint64(1)
		totalOut++
		tx.Transfers[i].Address, err = factom.NewFAAddress(list[i].Address)
		if err != nil {
			panic(fmt.Sprintf("%s is not a valid payout adress: %s", list[i].Address, err.Error()))
		}
	}
	batch.Transactions = append(batch.Transactions, tx)

	for {
		if len(batch.Transactions) == 0 {
			os.Exit(1)
		}
		entry, _ := batch.Sign(fs)
		data, _ := entry.MarshalBinary()
		if len(data) > 10275 {
			batch.Transactions[0].Transfers = batch.Transactions[0].Transfers[1:]
			continue
		}

		txid, err := entry.ComposeCreate(nil, factom.NewClient(), es)
		if err != nil {
			if strings.Contains(err.Error(), "length exceeds 10275") {
				batch.Transactions[0].Transfers = batch.Transactions[0].Transfers[1:]
				continue
			}
			panic(fmt.Errorf("unable to submit entry: %s", err.Error()))
		}

		fmt.Println(entry.Hash.String(), txid.String())
		fmt.Println(len(batch.Transactions[0].Transfers))
		break
	}

}

func globalRich(cl *srv.Client, count int) []srv.ResultGlobalRichList {
	var params srv.ParamsGetGlobalRichList
	params.Count = count

	var res []srv.ResultGlobalRichList
	err := cl.Request("get-global-rich-list", params, &res)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return res
}
