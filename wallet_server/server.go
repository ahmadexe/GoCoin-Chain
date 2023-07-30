package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type WalletServer struct {
	port    uint16
	gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}


// The gateway where the blockchain server is running is passed as a parameter to the wallet server. 
// This is because the wallet server needs to know where to send the transaction request and fetch other details.
func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func (ws *WalletServer) Wallet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		wlt := wallet.NewWallet()
		m, _ := wlt.MarshalJSON()
		io.WriteString(w, string(m))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Method not allowed")
	}
}

func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var tr transaction.TransactionResponse
		err := decoder.Decode(&tr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Bad Request")
			return
		}
		if !tr.Validate() {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Bad Request")
			return
		}

		publicKey := utils.PublicKeyFromString(*tr.SenderPublicKey)
		privateKey := utils.PrivateKeyFromString(*tr.SenderPrivateKey, publicKey)
		w.Header().Add("Content-Type", "application/json")

		transaction := wallet.NewTransaction(privateKey, publicKey, *tr.SenderBlockchainAddress, *tr.RecipientBlockchainAddress, *tr.Value)

		signature := transaction.GenerateSignature()
		signatureStr := signature.String()

		bt := &block.TransactionRequest{
			SenderPublicKey:           tr.SenderPublicKey,
			SenderChainAddress:        tr.SenderBlockchainAddress,
			Signature:                 &signatureStr,
			RecepientChainhainAddress: tr.RecipientBlockchainAddress,
			Value:                     tr.Value}

		m, _ := json.Marshal(bt)
		buf := bytes.NewBuffer(m)
		fmt.Println("http://"+ws.Gateway()+"/transaction")
		resp, err := http.Post("http://"+ws.Gateway()+"/transactions", "application/json", buf)
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Internal Server Error 1")
			return
		}

		if resp.StatusCode == http.StatusCreated {
			io.WriteString(w, string(rune(http.StatusCreated)))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Internal Server Error 2")

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Method not allowed")
	}
}

func (ws *WalletServer) WalletAmount(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		chainAddress := r.URL.Query().Get("blockchain_address")
		endPoint := fmt.Sprintf("http://%s/amount?blockchain_address=%s", ws.Gateway(), chainAddress)
		client := &http.Client{}
		resp, err := client.Get(endPoint)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Internal Server Error")
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			decoder := json.NewDecoder(resp.Body)
			var ar blockchain.AmountResponse
			err := decoder.Decode(&ar)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println("Internal Server Error")
				return
			}
			w.Header().Add("Content-Type", "application/json")
			m, _ := ar.MarshalJSON()
			io.WriteString(w, string(m))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Internal Server Error")

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Method not allowed")
	}
}

func (ws *WalletServer) Start() {
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	http.HandleFunc("/wallet/amount", ws.WalletAmount)
	log.Printf("Wallet server listening on port %v\n", ws.Port())
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
