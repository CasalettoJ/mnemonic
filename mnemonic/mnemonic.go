package mnemonic

import (
	"crypto/rand"
	"fmt"
	"log"
)

// Run runs the mnemonic process (printing for debug)
func Run() {
	// Step 1: Create a random sequence of 256 bits
	randomBits := make([]byte, 64)
	_, err := rand.Read(randomBits)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Step 1: Create 256 Random bits: %x\n", randomBits)

	// Step 2:
}
