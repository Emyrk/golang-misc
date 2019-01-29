package main

import (
	"encoding/json"
	"fmt"
	"os"

	"io/ioutil"

	"github.com/FactomProject/factomd/common/messages/msgsupport"
)

func main() {
	file, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, _ := ioutil.ReadAll(file)
	//v, err := hex.DecodeString(os.Args[1])
	//if err != nil {
	//	panic(err)
	//}

	msg, err := msgsupport.UnmarshalMessage(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(msg.String())
	str, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(str))
}
