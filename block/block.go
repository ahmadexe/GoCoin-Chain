package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ahmadexe/GoCoin-Chain/transaction"
)

type Block struct {
	Nonce        int                        `json:"nonce"`
	PreviousHash [32]byte                   `json:"previousHash"`
	Transactions []*transaction.Transaction `json:"transactions"`
	TimeStamp    int64                      `json:"timeStamp"`
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*transaction.Transaction) *Block {
	return &Block{nonce, previousHash, transactions, time.Now().UnixNano()}
}

func (b *Block) Print() {
	fmt.Printf("nonce: %d\n", b.Nonce)
	fmt.Printf("previousHash: %x\n", b.PreviousHash)
	fmt.Printf("timeStamp: %d\n", b.TimeStamp)
	for _, t := range b.Transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256(m)
}
