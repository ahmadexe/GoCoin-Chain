package main

import (
	"fmt"

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

	chainAddress := "my chain address"
	blockchain := blockchain.NewBlockchain(chainAddress)
	blockchain.AddTransaction("A", "B", 1.0)
	blockchain.Mining()
	blockchain.Print()

	blockchain.AddTransaction("C", "D", 2.0)
	blockchain.AddTransaction("X", "Y", 3.0)
	blockchain.Mining()
	blockchain.Print()

	fmt.Println("A's balance is", blockchain.CalculataBalance("A"))
	fmt.Println("B's balance is", blockchain.CalculataBalance("B"))
	fmt.Println("C's balance is", blockchain.CalculataBalance("C"))
	fmt.Println("D's balance is", blockchain.CalculataBalance("D"))
	fmt.Println("X's balance is", blockchain.CalculataBalance("X"))
	fmt.Println("Y's balance is", blockchain.CalculataBalance("Y"))
	fmt.Println("my chain address's balance is", blockchain.CalculataBalance(chainAddress))
}
