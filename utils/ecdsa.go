package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func (s *Signature) String() string {
	return fmt.Sprintf("%064x%064x", s.R, s.S)
}

func SignatureFromString(str string) *Signature {
	x, y := StringToBigIntTuple(str)
	return &Signature{R: &x, S: &y}
}

func StringToBigIntTuple(str string) (big.Int, big.Int) {
	bx, _ := hex.DecodeString(str[:64])
	by, _ := hex.DecodeString(str[64:])

	x := big.Int{}
	_ = x.SetBytes(bx)

	y := big.Int{}
	_ = y.SetBytes(by)

	return x, y
}

func PublicKeyFromString(str string) *ecdsa.PublicKey {
	x, y := StringToBigIntTuple(str)
	return &ecdsa.PublicKey{Curve: elliptic.P256(), X: &x, Y: &y}
}

func PrivateKeyFromString(str string, publicKey *ecdsa.PublicKey) *ecdsa.PrivateKey {
	x, _ := hex.DecodeString(str)
	var bi big.Int
	bi.SetBytes(x)
	return &ecdsa.PrivateKey{PublicKey: *publicKey, D: &bi}
}
