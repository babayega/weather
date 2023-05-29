package internal

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Chain struct {
	client *ethclient.Client
}

func NewChain() *Chain {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	return &Chain{
		client: client,
	}
}

// TODO: Contract call to check whether the address is registered or not.
func (c *Chain) AddressRegistered(address string) bool {
	return true
}
