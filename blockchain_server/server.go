package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ahmadexe/GoCoin-Chain/blockchain"
	"github.com/ahmadexe/GoCoin-Chain/wallet"
)


type BlockchainServer struct {
	port uint16
}

var cache map[string]*blockchain.Blockchain = make(map[string]*blockchain.Blockchain)

func (bcs *BlockchainServer) GetBlockchain() *blockchain.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minersWallter := wallet.NewWallet()
		bc = blockchain.NewBlockchain(minersWallter.BlockchainAddress, bcs.Port())
		cache["blockchain"] = bc
		
		// ! Don't log private keys in a real application
		log.Println("Created a new blockchain")
		log.Printf("Private key: %v\n", minersWallter.PrivateKey)
		log.Printf("Public key: %v\n", minersWallter.PrivateKey)
		log.Printf("Blockchain Address key: %v\n", minersWallter.PrivateKey)
	}
	return bc
}

func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
