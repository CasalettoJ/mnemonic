package mnemonic

import (
	"crypto/sha512"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	saltPrepend = "mnemonic"
	rounds      = 2048
	seedBytes   = 64
)

// GenerateSeed runs the pbkdf2 key stretching function with passphrase and salt to generate a 512bit key
func GenerateSeed(words []string, passphrase string) []byte {
	bits := pbkdf2.Key([]byte(strings.Join(words, " ")), []byte(strings.Join([]string{saltPrepend, passphrase}, "")), rounds, seedBytes, sha512.New)
	return bits
}
