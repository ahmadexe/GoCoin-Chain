package block

import (
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	transactions []string
	timeStamp    int64
}

func NewBlock(nonce int, previousHash string, transactions []string) *Block {
	return &Block{nonce, previousHash, transactions, time.Now().UnixNano()}
}

func (b *Block) Print() {
	fmt.Printf("nonce: %d\n", b.nonce)
	fmt.Printf("previousHash: %s\n", b.previousHash)
	fmt.Printf("transactions: %s\n", b.transactions)
	fmt.Printf("timeStamp: %d\n", b.timeStamp)
}


