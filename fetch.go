package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thadev404/smart_contract/contracts"
)

const address string = "0x1c129f46e1b512fae051959f5692126e51ef8279"

func main() {
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial("https://rinkeby.infura.io/v3/f5a4e8b0989e4ec79ca246113ad7309a")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Create a new instance of the Inbox contract bound to a specific deployed contract
	contract, err := contracts.NewInbox(common.HexToAddress(address), blockchain)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n", address)
	}

	fmt.Println(contract.Message(nil))

}
