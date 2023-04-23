package main

import (
	"fmt"

	"golang.org/x/crypto/blake2b"
)

func main() {
	text := "GeorgePradap"
	textBytes := []byte(text)

	// Create a Blake2b hash with a 32-byte output size
	hash, err := blake2b.New256(nil)
	if err != nil {
		panic(err)
	}

	// Write the input text bytes to the hash
	hash.Write(textBytes)

	// Get the final hash result as a 32-byte slice
	hash32 := hash.Sum(nil)

	fmt.Println("Original Text:", text)
	fmt.Println("32-byte Hash:", len(hash32))
}
