package main

import (
	"fmt"

	"github.com/ahmadexe/GoCoin-Chain/blockchain"
	"github.com/ahmadexe/GoCoin-Chain/wallet"
)

func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	t := wallet.NewTransaction(walletA.PrivateKey, walletA.PublicKey, walletA.BlockchainAddress, walletB.BlockchainAddress, 1.0)

	blockchain := blockchain.NewBlockchain(walletM.BlockchainAddress)
	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress, walletB.BlockchainAddress, 1.0, walletA.PublicKey, t.GenerateSignature())
	fmt.Println("Added:=",isAdded)
	blockchain.Mining()

	blockchain.Print()

	fmt.Printf("Balance of %s: %f\n", walletA.BlockchainAddress, blockchain.CalculateBalance(walletA.BlockchainAddress))
	fmt.Printf("Balance of %s: %f\n", walletB.BlockchainAddress, blockchain.CalculateBalance(walletB.BlockchainAddress))
	fmt.Printf("Balance of %s: %f\n", walletM.BlockchainAddress, blockchain.CalculateBalance(walletM.BlockchainAddress))



}
