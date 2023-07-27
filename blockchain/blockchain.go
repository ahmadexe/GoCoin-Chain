package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ahmadexe/GoCoin-Chain/block"
	"github.com/ahmadexe/GoCoin-Chain/transaction"
	"github.com/ahmadexe/GoCoin-Chain/utils"
)

type Blockchain struct {
	TransactionPool   []*transaction.Transaction
	Chain             []*block.Block
	BlockchainAddress string
	Port              uint16
}

func (bc *Blockchain) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TransactionPool   []*transaction.Transaction `json:"transactionPool"`
		Chain             []*block.Block             `json:"chain"`
		BlockchainAddress string                     `json:"blockchainAddress"`
		Port              uint16                     `json:"port"`
	}{
		TransactionPool:   bc.TransactionPool,
		Chain:             bc.Chain,
		BlockchainAddress: bc.BlockchainAddress,
		Port:              bc.Port,
	})
}

const (
	MINING_DIFFICULTY = 3
	MINING_REWARD     = 1.0
	MINING_SENDER     = "THE BLOCKCHAIN"
)

func NewBlockchain(blockchainAddress string, port uint16) *Blockchain {
	b := &block.Block{}
	bc := &Blockchain{
		[]*transaction.Transaction{},
		[]*block.Block{},
		blockchainAddress,
		port,
	}
	bc.createBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) createBlock(nonce int, previousHash [32]byte) *block.Block {
	block := block.NewBlock(nonce, previousHash, bc.TransactionPool)
	bc.Chain = append(bc.Chain, block)
	bc.TransactionPool = []*transaction.Transaction{}
	return block
}

func (bc *Blockchain) AddTransaction(senderChainAddress string, recipientChainAddress string, value float32, senderPublicKey *ecdsa.PublicKey, signature *utils.Signature) bool {

	transaction := transaction.NewTransaction(senderChainAddress, recipientChainAddress, value)

	if senderChainAddress == MINING_SENDER {
		bc.TransactionPool = append(bc.TransactionPool, transaction)
		return true
	} else if bc.verifyTransaction(senderPublicKey, signature, transaction) {
		bc.TransactionPool = append(bc.TransactionPool, transaction)
		return true
	}
	return false
}

func (bc *Blockchain) CreateTransaction(senderChainAddress string, recipientChainAddress string, value float32, senderPublicKey *ecdsa.PublicKey, signature *utils.Signature) bool {
	isTransacted := bc.AddTransaction(senderChainAddress, recipientChainAddress, value, senderPublicKey, signature)
	
	// TODO: Sync Blockchain servers

	return isTransacted
}

func (bc *Blockchain) LastBlock() *block.Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) CopyTransactions() []*transaction.Transaction {
	var transactions []*transaction.Transaction
	for _, t := range bc.TransactionPool {
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
	guessBlock := block.Block{TimeStamp: 0, Nonce: nonce, PreviousHash: previousHash, Transactions: transactions}
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
	bc.AddTransaction(MINING_SENDER, bc.BlockchainAddress, MINING_REWARD, nil, nil)
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.createBlock(nonce, previousHash)
	fmt.Println("Mining is successful!")
	return true
}

func (bc *Blockchain) CalculateBalance(address string) float32 {
	var total float32 = 0.0
	for _, b := range bc.Chain {
		for _, t := range b.Transactions {
			value := t.Value
			if address == t.RecipientChainAddress {
				total += value
			} else if address == t.SenderChainAddress {
				total -= value
			}
		}
	}
	return total
}

func (bc *Blockchain) verifyTransaction(senderPublicKey *ecdsa.PublicKey, sig *utils.Signature, t *transaction.Transaction) bool {
	m, _ := json.Marshal(t)
	hash := sha256.Sum256(m)
	return ecdsa.Verify(senderPublicKey, hash[:], sig.R, sig.S)
}

func (bc *Blockchain) Print() {
	for i, block := range bc.Chain {
		fmt.Println(strings.Repeat("-", 25), i, strings.Repeat("-", 25))
		block.Print()
		fmt.Println(strings.Repeat("-", 53))
	}
}


