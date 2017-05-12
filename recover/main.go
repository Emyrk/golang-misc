package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/FactomProject/bolt"
	fw "github.com/FactomProject/factoid/wallet"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/factom/wallet"
	"github.com/FactomProject/factomd/database/boltdb"
)

var _ = hex.Decode
var _ = wallet.ApiVersion

type Raw struct {
	data []byte
}

func (r *Raw) MarshalBinary() ([]byte, error) {
	return r.data, nil
}

func (r *Raw) UnmarshalBinary(data []byte) error {
	r.data = data
	return nil
}

func (r *Raw) UnmarshalBinaryData(data []byte) ([]byte, error) {
	r.data = data
	return nil, nil
}

type WalletEntryWrapper struct {
	WE *fw.WalletEntry
}

func (wew *WalletEntryWrapper) MarshalBinary() ([]byte, error) {
	return wew.WE.MarshalBinary()
}

func (wew *WalletEntryWrapper) UnmarshalBinary(data []byte) error {
	return wew.WE.UnmarshalBinary(data)
}

func (wew *WalletEntryWrapper) UnmarshalBinaryData(data []byte) ([]byte, error) {
	return wew.WE.UnmarshalBinaryData(data)
}

// OpenBoltDB opens a boltDB if it exists
func OpenBoltDB(boltPath string) (db *boltdb.BoltDB, reterr error) {
	// check if the file exists or if it is a directory
	fileInfo, err := os.Stat(boltPath)
	if err == nil {
		if fileInfo.IsDir() {
			return nil, fmt.Errorf("The path %s is a directory.  Please specify a file name.", boltPath)
		}
	}

	if os.IsNotExist(err) {
		return nil, fmt.Errorf("No file exists at %s", boltPath)
	}

	if err != nil && !os.IsNotExist(err) { //some other error, besides the file not existing
		fmt.Printf("database error %s\n", err)
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Could not use wallet file \"%s\"\n%v\n", boltPath, r)
			fmt.Println("Trying another method....")
			tdb, err := bolt.Open(boltPath, 0600, nil)
			if err != nil {
				fmt.Println("Database could not be opend", err.Error())
			} else {
				fmt.Println("Database opened, but we will close as we cannot progress.")
				tdb.Close()
			}
			db = nil
			reterr = err
			return
		}
	}()

	db = boltdb.NewBoltDB(nil, boltPath)
	return db, nil
}

func main() {
	log.SetOutput(os.Stderr)
	var (
		// Optionally find files in other directory
		dir   = flag.String("dir", ".", "Path to directory")
		brute = flag.Bool("b", false, "Brute force file")
	)

	flag.Parse()
	os.Remove("new.db")

	path := *dir
	if *dir == "." {
		path = ""
	}

	// Find all files in this directory
	files, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Fatalf("Had an issue reading the directory: %s\n", err.Error())
	}

	var dbfiles []os.FileInfo
	for _, f := range files {
		if strings.Contains(path+f.Name(), ".db") {
			dbfiles = append(dbfiles, f)
		}
	}

	fmt.Printf("Found %d db files...\n", len(dbfiles))
	for _, dbf := range dbfiles {
		fmt.Printf("  |- %s\n", dbf.Name())
	}

	// Go through files
	for _, f := range dbfiles {
		filename := path + f.Name()
		if *brute {
			bruteopen(filename, f.Name())
		} else {
			openRawBolt(filename)
		}
	}
}

func bruteopen(filename string, name string) {
	fmt.Println("Brute opening: ", filename)
	f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("Error opening %s: %s\n", filename, err.Error())
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("Error Reading %s: %s\n", filename, err.Error())
	}

	start := time.Now()
	all := len(data) - 31
	log.Printf("%d bytes to parse through, meaning %d addresses to generate", len(data), all)
	tot := 0
	n := time.Now()
	for i := range data {
		if i+32 > len(data) {
			break
		}
		fa, err := factom.MakeFactoidAddress(data[i : i+32])
		if err != nil {
			fmt.Println("Cannot Make Address:", err.Error())
			continue
		}
		tot++
		fmt.Println(fa.String(), "::", fa.SecString())
		if tot%100 == 0 {
			s := time.Since(n).Seconds()
			log.Printf("%d/%d addresses generated. %f/s", tot, all, 100/s)
			n = time.Now()
		}
	}
	s := time.Since(start).Seconds()
	log.Printf("%d/%d addresses generated. Took %f seconds", tot, all, s)
}

func openRawBolt(filename string) {
	// openFile(path, f)
	db, err := OpenBoltDB(filename)
	if err != nil {
		fmt.Printf("Error with %s: %s\n", filename, err.Error())
		return
	}

	fmt.Println("Successfully opened", filename)
	buckets, err := db.ListAllBuckets()
	if err != nil {
		fmt.Println("Error getting buckets:", err.Error())
		return
	}

	for _, b := range buckets {
		keys, err := db.ListAllKeys(b)
		if err != nil {
			fmt.Printf("Error getting keys for bucket %x :: %s\n", b, string(b))
			return
		}
		for _, k := range keys {
			fa, err := factom.MakeFactoidAddress(k)
			if err == nil {
				fmt.Printf("From Key: %s :: %s\n", fa.String(), fa.SecString())
			}

			raw := new(Raw)
			rI, err := db.Get(b, k, raw)
			if err != nil {
				fmt.Printf("Error getting value for key %x :: %s\n", k, string(k))
				return
			}

			r := rI.(*Raw)
			if len(r.data) > 32 {
				sec := r.data[len(r.data)-64 : len(r.data)-32]
				fa, err := factom.MakeFactoidAddress(sec)
				if err == nil {
					fmt.Printf("From Value: %s :: %s\n", fa.String(), fa.SecString())
				}
			}
		}
	}
}

// openFile opens a file as a wallet
func openFile(path string, f os.FileInfo) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic with file %s: %v\n", f.Name(), r)
		}
	}()
	// Only open *.db files
	if strings.Contains(f.Name(), ".db") {
		if _, err := os.Stat(f.Name()); err == nil {
		} else {
			// file does not exist, just skipt it
			return
		}

		// Import V1 wallet
		wal, err := wallet.ImportV1Wallet(f.Name(), "new.db")
		if err != nil {
			log.Printf("%s could not be opened: %s\n", f.Name(), err.Error())
		} else {
			fas, _, err := wal.GetAllAddresses()
			if err != nil {
				log.Printf("%s was opened, but addresses could not be retrieved: %s\n", f.Name(), err.Error())
			} else {
				fmt.Printf("- %s has %d factoid addresses\n", f.Name(), len(fas))
				for _, fa := range fas {
					fmt.Printf("-- Factoid Address from %s:\n  %s\n  %s\n", f.Name(), fa.SecString(), fa.String())
				}
			}
			// Clean up files created
			wal.Close()
			os.Remove("new.db")
		}
	}
}
