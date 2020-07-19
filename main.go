package main

import (
	"barcode-keyboard-emulator/app"
	"flag"
	"log"
)

func main() {
	// Set flag
	// Default ":9090" - use all host ip and port 9090
	addr := flag.String("addr", "", "Set used address")
	port := flag.String("port", "9090", "Set used port")
	flag.Parse()

	// Start server
	err := app.Run(Version, *addr, *port)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
