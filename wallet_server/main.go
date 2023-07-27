package main

import "flag"

func main() {
	port := flag.Uint("port", 5050, "TCP Port Number for Wallet Server")
	gateway := flag.String("gateway", "0.0.0.0:8080", "Gateway URL for Blockchain Server")
	flag.Parse()
	app := NewWalletServer(uint16(*port), *gateway)
	app.Start()
}
