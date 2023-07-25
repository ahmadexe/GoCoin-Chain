package utils

import (
	"encoding/hex"
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
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
