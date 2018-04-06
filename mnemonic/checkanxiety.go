package mnemonic

import "log"

// CheckAnxiety checks if we gotta panic.
func checkAnxiety(err error) {
	if err != nil {
		log.Panic(err)
	}
}
