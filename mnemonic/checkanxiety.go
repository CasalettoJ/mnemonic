package mnemonic

import "log"

// CheckAnxiety checks if we gotta panic.
func CheckAnxiety(err error) {
	if err != nil {
		log.Panic(err)
	}
}
