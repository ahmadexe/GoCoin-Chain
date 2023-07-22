package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/ahmadexe/GoCoin-Chain/utils"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

/*
* Steps to follow:
* 1. Generate a private key
* 2. Generate a public key from the private key
* 3. Perform SHA256 on the public key
* 4. Perform RIPEMD160 on the result of SHA256
* 5. Add version byte in front of RIPEMD160 hash (0x00 for Main Network)
* 6. Perform SHA256 on the extended RIPEMD160 result
* 7. Perform SHA256 on the result of the previous SHA256
* 8. Take the first 4 bytes of the second SHA256 hash. This is the address checksum
* 9. Add the 4 checksum bytes from stage 7 at the end of extended RIPEMD160 hash from stage 5.
* 10. Convert the result from a byte string into a base58 string using Base58Check encoding. This is the most commonly used Bitcoin Address format
 */

type Wallet struct {
	PrivateKey        *ecdsa.PrivateKey
	PublicKey         *ecdsa.PublicKey
	BlockchainAddress string
}

func NewWallet() *Wallet {
	w := &Wallet{}
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.PrivateKey = privateKey
	w.PublicKey = &w.PrivateKey.PublicKey

	// Step 3
	h2 := sha256.New()
	h2.Write(w.PublicKey.X.Bytes())
	h2.Write(w.PublicKey.Y.Bytes())
	digest2 := h2.Sum(nil)

	// Step 4
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)

	// Step 5
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3)

	// Step 6
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)

	// Step 7
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)

	// Step 8
	chksum := digest6[:4]

	// Step 9
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], chksum[:])

	// Step 10
	address := base58.Encode(dc8)
	w.BlockchainAddress = address

	return w
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.PrivateKey.D.Bytes())
}

func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.PublicKey.X.Bytes(), w.PublicKey.Y.Bytes())
}

type Transaction struct {
	senderPrivateKey      *ecdsa.PrivateKey
	senderPublicKey       *ecdsa.PublicKey
	senderChainAddress    string
	recipientChainAddress string
	value                 float32
}

func NewTransaction(senderPrivateKey *ecdsa.PrivateKey, senderPublicKey *ecdsa.PublicKey, senderChainAddress string, recipientChainAddress string, value float32) *Transaction {
	return &Transaction{senderPrivateKey, senderPublicKey, senderChainAddress, recipientChainAddress, value}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderChainAddress    string  `json:"senderChainAddress"`
		RecipientChainAddress string  `json:"recipientChainAddress"`
		Value                 float32 `json:"value"`
	}{
		SenderChainAddress:    t.senderChainAddress,
		RecipientChainAddress: t.recipientChainAddress,
		Value:                 t.value,
	})
}

func (t *Transaction) GenerateSignature() *utils.Signature {
	m, _ := json.Marshal(t)
	h := sha256.Sum256([]byte(m))
	r, s, _ := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
	return &utils.Signature{R: r, S: s}
}

func (w *Wallet) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PrivateKey        string `json:"privateKey"`
		PublicKey         string `json:"publicKey"`
		BlockchainAddress string `json:"blockchainAddress"`
	}{
		PrivateKey:        w.PrivateKeyStr(),
		PublicKey:         w.PublicKeyStr(),
		BlockchainAddress: w.BlockchainAddress,
	})
}
