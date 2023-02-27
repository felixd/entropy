package main

import (
	"crypto/rand"
	"encoding/binary"
	"os"
	"flag"
)

func main() {

	// Parse the command-line arguments. By default 8 milion bytes are generated
	var numBytes int
	flag.IntVar(&numBytes, "n", 8*1000000, "number of bytes to be generated")
	flag.Parse()

	// Generate 8 million random numbers using crypto/rand
	randBytes := make([]byte, numBytes)
	if _, err := rand.Read(randBytes); err != nil {
		panic(err)
	}

	// Convert the slice of random numbers to a binary file
	f, err := os.Create("random.bin")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := binary.Write(f, binary.LittleEndian, randBytes); err != nil {
		panic(err)
	}
}
