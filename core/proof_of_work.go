package core

import (
	"bytes"
	_ "strconv"
)

const Difficulty = 3

// MineBlock PoW performs a proof of work, adjusting nonce until the hash starts with a number of zeroes
func (b *Block) MineBlock() {
	for !b.IsValidProof() {
		b.Nonce++
		b.Hash = b.CalculateHash()
	}
}

// IsValidProof checks if the block's hash satisfies the required difficulty (e.g., number of leading zeroes)
func (b *Block) IsValidProof() bool {
	hashBytes := []byte(b.Hash)
	return bytes.HasPrefix(hashBytes, bytes.Repeat([]byte{0}, Difficulty))
}
