package main

import (
	"fmt"

	"github.com/umbracle/ethgo/jsonrpc"
)

func main() {

	cient, err := jsonrpc.NewClient("https://mainnet.infura.io")
	if err != nil {

		panic(err)

	}
	fmt.Print(cient)
}
