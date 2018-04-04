// mnemonic generates a 24 mnemonic word string following
// https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki
// https://github.com/bitcoinbook/bitcoinbook/blob/develop/ch05.asciidoc
package mnemonic

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"strings"
)

const (
	wordCount       = 24
	bitSplitAmount  = 11
	randomBitAmount = 256
)

func generateRandomBits() []byte {
	randomBits := make([]byte, 256/8)
	_, err := rand.Read(randomBits)
	CheckAnxiety(err)
	return randomBits
}

func generateChecksum(randomBits []byte) ([]byte, byte) {
	hash := sha256.Sum256(randomBits)
	return hash[:], hash[0]
}

func splitBits(randomBits []byte) []string {
	var bytesToBits []string
	var bitSplits []string
	for _, currentByte := range randomBits {
		bytesToBits = append(bytesToBits, fmt.Sprintf("%08b", currentByte))
	}
	bits := strings.Join(bytesToBits, "")
	bitRunes := []rune(bits)
	for i := 0; i < wordCount; i++ {
		bitString := string(bitRunes[bitSplitAmount*i : (bitSplitAmount*i)+bitSplitAmount])
		bitSplits = append(bitSplits, bitString)
	}
	return bitSplits
}

// Run runs the mnemonic process (printing for debug)
func Run() {
	// Step 1: Create a random sequence of 256 bits
	randomBits := generateRandomBits()
	fmt.Printf("Step 1: Create 256 Random bits: %x\n\n", randomBits)

	// Step 2: Create a checksum of the random sequence by taking the first (entropy-length/32)
	// bits of its SHA256 hash.
	fullHash, checksum := generateChecksum(randomBits)
	fmt.Printf("Step 2: Generate a checksum:\n%x\nand take the first (bits/32) bits:\n%x\n\n", fullHash, checksum)
	// Step 3: add it to the end of the generated bits
	randomBits = append(randomBits, checksum)
	fmt.Printf("Step 3: Add the checksum to the end of the generated bits:\n%x\n\n", randomBits)
	// Step 4: split the result into 11-bit-length segments
	separatedBits := splitBits(randomBits)
	fmt.Println("Step 4: split the result into 11-bit-length segments")
	for i, bitString := range separatedBits {
		if i != 0 {
			fmt.Printf(" | ")
		}
		fmt.Printf("%s", bitString)
	}
	fmt.Printf("\n\n")
}
