package structures

import "github.com/Suad0/blockchainGO/core"

type TransactionPool struct {
	Transactions []core.Transaction
}

// AddTransaction adds a transaction to the pool
func (tp *TransactionPool) AddTransaction(tx core.Transaction) {
	tp.Transactions = append(tp.Transactions, tx)
}
