package main

import (
	"fmt"

	"github.com/dant89/coinbase-go-v2"
)

func main() {

	c := coinbase.ApiKeyClient("BgpUvkgK0suj9apk", "W5lYiC5EWkXLAlqY4FJLec7Ma0ydRfZj")
	acc, _, err := c.GetAccounts(map[string]string{})

	//u, err := c.GetUser()
	if err != nil {
		panic(err)
	}
	fmt.Println(acc)
	//
	//f, err := c.GetBalance()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(f)
}
