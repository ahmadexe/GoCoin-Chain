package blockchain

import (
	"fmt"
	"strings"

	"github.com/ahmadexe/GoCoin-Chain/block"
)

type Blockchain struct {
	transactionPool []string
	chain           []*block.Block
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *block.Block {
	block := block.NewBlock(nonce, previousHash, []string{})
	bc.chain = append(bc.chain, block)
	return block
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{
		[]string{},
		[]*block.Block{},
	}
	bc.CreateBlock(0, "Init Hash")
	return bc
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Println(strings.Repeat("-", 25), i, strings.Repeat("-", 25))
		block.Print()
		fmt.Println(strings.Repeat("-", 53))
	}
}
