package blockchain

import (
	"fmt"
	"strings"

	"github.com/ahmadexe/GoCoin-Chain/block"
	"github.com/ahmadexe/GoCoin-Chain/transaction"
)

type Blockchain struct {
	transactionPool []*transaction.Transaction
	chain           []*block.Block
}

func NewBlockchain() *Blockchain {
	b := &block.Block{}
	bc := &Blockchain{
		[]*transaction.Transaction{},
		[]*block.Block{},
	}
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *block.Block {
	block := block.NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, block)
	bc.transactionPool = []*transaction.Transaction{}
	return block
}

func (bc *Blockchain) AddTransaction(senderChainAddress string, recipientChainAddress string, value float32) {
	transaction := transaction.NewTransaction(senderChainAddress, recipientChainAddress, value)
	bc.transactionPool = append(bc.transactionPool, transaction)
}

func (bc *Blockchain) LastBlock() *block.Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Println(strings.Repeat("-", 25), i, strings.Repeat("-", 25))
		block.Print()
		fmt.Println(strings.Repeat("-", 53))
	}
}