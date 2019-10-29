package main

import (
	"encoding/binary"
	"log"
	"os"
	qr "github.com/skip2/go-qrcode"
)

func main() {
	png, err := qr.Encode(string(os.Args[1]), qr.Medium, 256)
	if err != nil {
		log.Fatalf("Unable to read standard input: %s", err.Error())
	}
	binary.Write(os.Stdout, binary.LittleEndian, png)
}
