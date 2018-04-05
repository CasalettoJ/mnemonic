// Package mnemonic generates a 24 mnemonic word string following
// https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki
// https://github.com/bitcoinbook/bitcoinbook/blob/develop/ch05.asciidoc
package mnemonic

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

const (
	wordCount       = 24
	bitSplitAmount  = 11
	randomBitAmount = 256
)

func generateRandomBits() []byte {
	randomBits := make([]byte, randomBitAmount/8)
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

func mapDictionary(separatedBits []string) []string {
	var words []string
	for _, bits := range separatedBits {
		index, err := strconv.ParseInt(bits, 2, 64)
		CheckAnxiety(err)
		words = append(words, EnglishDict[index])
	}
	return words
}

// GenerateWords creates a set of 24 word mnemonic code
func GenerateWords() []string {
	randomBits := generateRandomBits()
	_, checksum := generateChecksum(randomBits)
	randomBits = append(randomBits, checksum)
	separatedBits := splitBits(randomBits)
	words := mapDictionary(separatedBits)
	return words
}
