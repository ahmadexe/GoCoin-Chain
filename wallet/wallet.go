package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

func NewWallet() *Wallet {
	w := &Wallet{}
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.PrivateKey = privateKey
	w.PublicKey = &privateKey.PublicKey
	return w
}

func (w *Wallet) PrivateKeyStr() string {
	return fmt.Sprintf("%x", w.PrivateKey.D.Bytes())
}

func (w *Wallet) PublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.PublicKey.X.Bytes(), w.PublicKey.Y.Bytes())
}
