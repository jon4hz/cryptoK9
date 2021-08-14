package mnemonic

import (
	"github.com/tyler-smith/go-bip39"
)

func IsValid(mnemonic string) bool {
	return bip39.IsMnemonicValid(mnemonic)
}
