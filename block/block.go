package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	Nonce        int      `json:"nonce"`
	PreviousHash [32]byte   `json:"previousHash"`
	Transactions []string `json:"transactions"`
	TimeStamp    int64    `json:"timeStamp"`
}

func NewBlock(nonce int, previousHash [32]byte, transactions []string) *Block {
	return &Block{nonce, previousHash, transactions, time.Now().UnixNano()}
}

func (b *Block) Print() {
	fmt.Printf("nonce: %d\n", b.Nonce)
	fmt.Printf("previousHash: %s\n", b.PreviousHash)
	fmt.Printf("transactions: %s\n", b.Transactions)
	fmt.Printf("timeStamp: %d\n", b.TimeStamp)
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256(m)
}
