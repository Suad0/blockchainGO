package miner

import "github.com/Suad0/blockchainGO/core"

type Miner struct {
	Blockchain *core.Blockchain
}

// Mine mines a new block
func (m *Miner) Mine(transactions []core.Transaction) {
	m.Blockchain.AddBlock(transactions)
}
