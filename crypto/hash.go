package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	// str := "George Pradap Mariya Raj"

	// // Create a new SHA-256 hash object
	// hashAlgo := sha256.New()

	// // Write the data to the hash object
	// hashAlgo.Write([]byte(str))

	// // Get the hash value
	// buffer := hashAlgo.Sum(nil)
	// fmt.Printf("buffer: %v\n", buffer)

	// // Convert the hash value to a string
	// hash := hex.EncodeToString(buffer)

	// fmt.Printf("hash : %v\n", hash)

	fmt.Printf("%x\n", sha256.Sum256([]byte("George Pradap Mariya Raj")))

	fmt.Printf("%x\n", sha1.Sum([]byte("George Pradap Mariya Raj")))

	fmt.Printf("%x\n", sha512.Sum512([]byte("George Pradap Mariya Raj")))

	fmt.Printf("%x\n", md5.Sum([]byte("George Pradap Mariya Raj")))

	
}
