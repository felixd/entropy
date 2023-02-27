package main

import (
	"crypto/rand"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
)

func uniqueBytes(byteSlice []byte) []byte {
	keys := make(map[byte]bool)
	list := []byte{}
	for _, entry := range byteSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

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

	uniqueVals := uniqueBytes(randBytes)

	fmt.Println("uniqueBytes/ randBytes:")
	fmt.Println(uniqueVals)
	fmt.Println(randBytes)

	numUniqueVals := len(uniqueVals)
	// Calculate the number of bits of entropy
	numBitsOfEntropy := numUniqueVals * 8

	// Simple Entropy check:
	fmt.Printf("Number of bits of entropy (unique Bytes) in the []byte slice: %d\n", numBitsOfEntropy)

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
