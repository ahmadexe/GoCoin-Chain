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

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *block.Block {
	block := block.NewBlock(nonce, previousHash ,[]string{})
	bc.chain = append(bc.chain, block)
	return block
}

func NewBlockchain() *Blockchain {
	b := &block.Block{}
	bc := &Blockchain{
		[]string{},
		[]*block.Block{},
	}
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Println(strings.Repeat("-", 25), i, strings.Repeat("-", 25))
		block.Print()
		fmt.Println(strings.Repeat("-", 53))
	}
}

func (bc *Blockchain) LastBlock() *block.Block {
	return bc.chain[len(bc.chain)-1]
}