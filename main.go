package main

import (
	"github.com/ahmadexe/GoCoin-Chain/blockchain"
)

func main() {
	// bc := blockchain.NewBlockchain()
	// bc.Print()
	// bc.CreateBlock(1, "Hash 1")
	// bc.CreateBlock(2, "Hash 2")
	// bc.Print()

	// b := &block.Block{
	// 	Nonce:        1,
	// 	Transactions: []string{},
	// 	TimeStamp:    1,
	// }

	// fmt.Printf("%x\n",b.Hash())

	blockchain := blockchain.NewBlockchain()
	// blockchain.Print()
	
	blockchain.AddTransaction("A", "B", 1.0)
	prevHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(1, prevHash)
	// blockchain.Print()

	blockchain.AddTransaction("C", "D", 2.0)
	blockchain.AddTransaction("X", "Y", 3.0)
	prevHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(2, prevHash)
	blockchain.Print()
}
