package main

import (
	"github.com/ahmadexe/GoCoin-Chain/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	// bc.Print()
	bc.CreateBlock(1, "Hash 1")
	bc.CreateBlock(2, "Hash 2")
	bc.Print()
}
