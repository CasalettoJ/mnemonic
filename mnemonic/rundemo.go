package mnemonic

import "fmt"

// RunDemo runs the mnemonic process with explanations for each step for help
func RunDemo() {
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
	// Step 5: Map the binary numbers to array indices in the dictionary
	words := mapDictionary(separatedBits)

	fmt.Println("Step 5: Map the binary numbers to array indices in the dictionary")
	for _, word := range words {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n\n")

	// Step 6: Generate a 512bit seed from the words and salt
	bits := GenerateSeed(words, "password")
	fmt.Println("Step 6: Generate a 512bit seed from the words and salt")
	fmt.Printf("%x\n\n", bits)
}
