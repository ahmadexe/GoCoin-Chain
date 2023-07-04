package transaction

import (
	"fmt"
	"strings"
)

type Transaction struct {
	SenderChainAddress    string  `json:"senderChainAddress"`
	RecipientChainAddress string  `json:"recipientChainAddress"`
	Value                 float32 `json:"value"`
}

func NewTransaction(senderAddress string, recipientAddress string, val float32) *Transaction {
	return &Transaction{senderAddress, recipientAddress, val}
}

func (t *Transaction) Print() {
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("senderChainAddress: %s\n", t.SenderChainAddress)
	fmt.Printf("recipientChainAddress: %s\n", t.RecipientChainAddress)
	fmt.Printf("value: %f\n", t.Value)
	fmt.Println(strings.Repeat("-", 40))
}
