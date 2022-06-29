package main

import (
	"fmt"

	"github.com/umbracle/ethgo/jsonrpc"
)

//Eth : 이더리움 네트워크 엔드포인트.
//Net : 클라이언트 정보.
var ethet = "https://api.polygonscan.com"
var apikay = "apikety"

func main() {

	client, err := jsonrpc.NewClient("https://api.baobab.klaytn.net:8651")

	if err != nil {

		panic(err)

	}
	eth := client.Eth()
	//count (uint64): 연결된 피어의 수입니다.
	count, err := client.Net().PeerCount()
	accounts, err := client.Eth().Accounts()
	number, err := client.Eth().BlockNumber()
	price, err := client.Eth().GasPrice()
	fmt.Println(price)
	fmt.Println(number)
	fmt.Println(count)
	fmt.Print(client)
	fmt.Println(eth)
	fmt.Println(accounts)

}
