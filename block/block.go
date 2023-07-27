package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ahmadexe/GoCoin-Chain/transaction"
)


type Block struct {
	TimeStamp    int64
	Nonce        int
	PreviousHash [32]byte
	Transactions []*transaction.Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*transaction.Transaction) *Block {
	return &Block{Nonce: nonce, PreviousHash: previousHash, Transactions: transactions, TimeStamp: time.Now().UnixNano()}
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TimeStamp    int64                      `json:"timeStamp"`
		Nonce        int                        `json:"nonce"`
		PreviousHash string                   `json:"previousHash"`
		Transactions []*transaction.Transaction `json:"transactions"`
	}{
		TimeStamp:    b.TimeStamp,
		Nonce:        b.Nonce,
		PreviousHash: fmt.Sprintf("%x",b.PreviousHash),
		Transactions: b.Transactions,
	})
}

func (b *Block) Print() {
	fmt.Printf("Nonce: %d\n", b.Nonce)
	fmt.Printf("PreviousHash: %x\n", b.PreviousHash)
	fmt.Printf("TimeStamp: %d\n", b.TimeStamp)
	for _, t := range b.Transactions {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}

type TransactionRequest struct {
	SenderPublicKey           *string  `json:"senderPublicKey"`
	SenderChainAddress        *string  `json:"senderChainAddress"`
	Signature                 *string  `json:"signature"`
	RecepientChainhainAddress *string  `json:"recepientChainhainAddress"`
	Value                     *float32 `json:"value"`
}

func (tr *TransactionRequest) Validate() bool {
	return tr.SenderPublicKey != nil && tr.SenderChainAddress != nil && tr.Signature != nil && tr.RecepientChainhainAddress != nil && tr.Value != nil
}