package main

import "flag"

func main() {
	port := flag.Uint("port", 8080, "TCP Port Number for Wallet Server")
	gateway := flag.String("gateway", "http://localhost:8080", "Gateway URL for Blockchain Server")
	flag.Parse()
	app := NewWalletServer(uint16(*port), *gateway)
	app.Start()
}
