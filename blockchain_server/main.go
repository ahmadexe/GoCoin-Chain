package main

import (
	"flag"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	port := flag.Uint("port", 8080, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := NewBlockchainServer(uint16(*port))
	fmt.Printf("Starting Blockchain Server on Port %d\n", app.Port())
	app.Run()
}
