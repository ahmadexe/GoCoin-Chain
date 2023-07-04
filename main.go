package main

import (
	"fmt"

	"github.com/ahmadexe/GoCoin-Chain/block"
)

func main() {
	// bc := blockchain.NewBlockchain()
	// bc.Print()
	// bc.CreateBlock(1, "Hash 1")
	// bc.CreateBlock(2, "Hash 2")
	// bc.Print()

	b := &block.Block{
		Nonce:        1,
		Transactions: []string{},
		TimeStamp:    1,
	}

	fmt.Printf("%x\n",b.Hash())

}
