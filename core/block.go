package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Timestamp    time.Time
	Transactions []Transaction
	PrevHash     string
	Hash         string
	Nonce        int
}

func (b *Block) CalculateHash() string {
	record := b.Timestamp.String() + b.PrevHash + string(rune(b.Nonce))
	hash := sha256.New()
	hash.Write([]byte(record))
	return hex.EncodeToString(hash.Sum(nil))
}

func NewBlock(transactions []Transaction, prevHash string) *Block {
	block := &Block{
		Timestamp:    time.Now(),
		Transactions: transactions,
		PrevHash:     prevHash,
		Nonce:        0,
	}
	block.Hash = block.CalculateHash()
	return block
}
