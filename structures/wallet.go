package structures

import (
	"crypto/rsa"
	"github.com/Suad0/blockchainGO/crypto"
)

type Wallet struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewWallet() *Wallet {
	privateKey, publicKey := crypto.GenerateKeyPair(2048)
	return &Wallet{PrivateKey: privateKey, PublicKey: publicKey}
}
