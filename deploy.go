package main

import (
	"fmt"
	"log"
	_ "os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thadev404/smart_contract/contracts"
)

const key = `{"address":"9e283dd051ab9b8c5053ee746ccf6104d2261d8e","crypto":{"cipher":"aes-128-ctr","ciphertext":"0587d5d4b3802e1aba158189063c4c7c288be009f35c09b30f80363fe15b9c62","cipherparams":{"iv":"46ae07bb9905fba304328723f925571c"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"324e93f9444994dad70f25fca21e0396825da94c2cf23f181fdad5620c22a5b1"},"mac":"3bb7811b5509754ddd2a8b445a70ae71bc87a4873f8f1f67316b7414a5e7e8c0"},"id":"93d20f29-cc09-4a02-9954-a4e18aa2616a","version":3}`

func mains() {
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial("https://rinkeby.infura.io/v3/f5a4e8b0989e4ec79ca246113ad7309a")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(key), "*******,")

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	address, _, _, _ := contracts.DeployInbox(
		auth,
		blockchain,
		"Hello World",
	)

	fmt.Printf("Contract pending deploy: 0x%x\n", address)
}
