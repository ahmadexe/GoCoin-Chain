package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ahmadexe/GoCoin-Chain/block"
	"github.com/ahmadexe/GoCoin-Chain/blockchain"
	"github.com/ahmadexe/GoCoin-Chain/transaction"
	"github.com/ahmadexe/GoCoin-Chain/utils"
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

func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		bc := bcs.GetBlockchain()
		m, _ := json.Marshal(bc)
		io.WriteString(w, string(m[:]))
	default:
		log.Println("Method not allowed")
	}
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

func (bcs *BlockchainServer) Transactions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()

		m, _ := json.Marshal(struct {
			Transactions []*transaction.Transaction `json:"transactions"`
			Length       int                        `json:"length"`
		}{
			Transactions: bc.TransactionPool,
			Length:       len(bc.TransactionPool),
		})

		io.WriteString(w, string(m[:]))

	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var t block.TransactionRequest
		err := decoder.Decode(&t)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Bad Request")
			return
		}
		if !t.Validate() {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Bad Request")
			return
		}

		publicKey := utils.PublicKeyFromString(*t.SenderPublicKey)
		signature := utils.SignatureFromString(*t.Signature)

		bc := bcs.GetBlockchain()

		isCreated := bc.CreateTransaction(*t.SenderChainAddress, *t.RecepientChainhainAddress, *t.Value, publicKey, signature)

		w.Header().Add("Content-Type", "application/json")

		if isCreated {
			w.WriteHeader(http.StatusCreated)
			log.Println("Transaction created")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Bad Request")
	}
}

func (bcs *BlockchainServer) Mine(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		bc := bcs.GetBlockchain()
		isMined := bc.Mining()
		if isMined {
			log.Println("Mined")
			w.WriteHeader(http.StatusCreated)
			return
		} else {
			log.Println("Not Mined")
			w.WriteHeader(http.StatusConflict)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Method not allowed")
	}
}

func (bcs *BlockchainServer) StartMine(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		bc := bcs.GetBlockchain()
		bc.StartMining()
		m := "Mining started"
		io.WriteString(w, string(m[:]))
		log.Println("Mining started")

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Method not allowed")
	}
}

func (bcs *BlockchainServer) Run() {

	http.HandleFunc("/", bcs.GetChain)
	http.HandleFunc("/transactions", bcs.Transactions)
	http.HandleFunc("/mine", bcs.Mine)
	http.HandleFunc("/mine/start", bcs.StartMine)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
