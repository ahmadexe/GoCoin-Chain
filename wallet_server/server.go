package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/ahmadexe/GoCoin-Chain/transaction"
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

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Method not allowed")
	}
}

func (ws *WalletServer) Start() {
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	log.Printf("Wallet server listening on port %v\n", ws.Port())
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))

}
