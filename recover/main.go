package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/FactomProject/bolt"
	fw "github.com/FactomProject/factoid/wallet"
	"github.com/FactomProject/factom/wallet"
	"github.com/FactomProject/factomd/database/boltdb"
)

var _ = wallet.ApiVersion

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
		}
	}()

	db = boltdb.NewBoltDB(nil, boltPath)
	return db, nil
}

func main() {
	var (
		// Optionally find files in other directory
		dir = flag.String("dir", ".", "Path to directory")
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
		// openFile(path, f)
		filename := path + f.Name()
		db, err := OpenBoltDB(filename)
		if err != nil {
			fmt.Printf("Error with %s: %s\n", filename, err.Error())
		}

		fmt.Println("Successfully opened", filename)
		buckets, err := db.ListAllBuckets()
		if err != nil {
			fmt.Println("Error getting buckets:", err.Error())
			continue
		}

		for _, b := range buckets {
			keys, err := db.ListAllKeys(b)
			if err != nil {
				fmt.Printf("Error getting keys for bucket %x :: %s\n", b, string(b))
				continue
			}
			for _, k := range keys {
				fmt.Println(len(k))
				we := new(WalletEntryWrapper)
				wewI, err := db.Get(b, k, we)
				if err != nil {
					fmt.Printf("Error getting value for key %x :: %s\n", k, string(k))
					continue
				}

				wew := wewI.(*WalletEntryWrapper)
				add, err := wew.WE.GetAddress()
				if err != nil {
					fmt.Printf("Error printing value for key %x :: %s\n", k, string(k))
					continue
				}
				fmt.Printf("WE: %s\n", add.String())
			}

		}

		var _ = db
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
