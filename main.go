package main

import (
	"fmt"

	"github.com/ahmadexe/GoCoin-Chain/block"
)

func main() {
	fmt.Println("Init")
	b := block.NewBlock(0, "", []string{})
	b.Print()
}
