package main

import (
	"fmt"

	"github.com/ahmadexe/GoCoin-Chain/wallet"
)

func main() {
	w := wallet.NewWallet()
	t := wallet.NewTransaction(w.PrivateKey, w.PublicKey, w.BlockchainAddress, "B", 1.0)
	fmt.Printf("Transaction Signature: %s\n", t.GenerateSignature())
}
