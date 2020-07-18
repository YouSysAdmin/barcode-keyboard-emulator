package main

import (
	"barcode-keyboard-emulator/app"
	"log"
)

func main() {
	// Start server
	err := app.Run()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
