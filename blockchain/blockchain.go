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
	blockchainAddress string
}

const (
	MINING_DIFFICULTY = 3
	MINING_REWARD = 1.0
	MINING_SENDER = "THE BLOCKCHAIN"
)

func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &block.Block{}
	bc := &Blockchain{
		[]*transaction.Transaction{},
		[]*block.Block{},
		blockchainAddress,
	}
	bc.createBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) createBlock(nonce int, previousHash [32]byte) *block.Block {
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

func (bc *Blockchain) CopyTransactions() []*transaction.Transaction {
	var transactions []*transaction.Transaction
	for _, t := range bc.transactionPool {
		transaction := &transaction.Transaction{
			SenderChainAddress:    t.SenderChainAddress,
			RecipientChainAddress: t.RecipientChainAddress,
			Value:                 t.Value,
		}

		transactions = append(transactions, transaction)
	}
	return transactions
}

func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*transaction.Transaction, difficulty int) bool {
	zeroes := strings.Repeat("0", difficulty)
	guessBlock := block.Block{Nonce: nonce, PreviousHash: previousHash, Transactions: transactions, TimeStamp: 0}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeroes
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactions()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce++
	}
	return nonce
}

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.createBlock(nonce, previousHash)
	fmt.Println("Mining is successful!")
	return true
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Println(strings.Repeat("-", 25), i, strings.Repeat("-", 25))
		block.Print()
		fmt.Println(strings.Repeat("-", 53))
	}
}
