package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Generate 1 million random numbers using crypto/rand
	randBytes := make([]byte, 8*1000000)
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
