package main

import "github.com/ahmadexe/GoCoin-Chain/wallet"

func main() {
	w := wallet.NewWallet()
	println(w.PrivateKeyStr())
	println(w.PublicKeyStr())
}
