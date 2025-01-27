package parser

import (
	"log"
	"tx-parser/models"
)

func (tp *TxParser) GetTransactions(address string) []models.Transaction {
	tp.mu.RLock()
	defer tp.mu.RUnlock()

	transactionCh, exists := tp.subscribedAddr[address]
	if !exists {
		log.Printf("Address not subscribed: %s", address)
		return nil
	}

	var transactions []models.Transaction
	for transaction := range transactionCh {
		transactions = append(transactions, transaction)
	}
	return transactions
}
