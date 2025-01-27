package parser

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"strconv"
	"tx-parser/models"
)

func (tp *TxParser) Subscribe(address string) bool {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	if _, exists := tp.subscribedAddr[address]; exists {
		log.Printf("address %s is already subscribed", address)
		return false
	}

	txChannel := make(chan models.Transaction, 10) // Buffered channel to avoid blocking
	tp.subscribedAddr[address] = txChannel

	go tp.listenForTransactions(address, txChannel)

	log.Printf("Subscribed to address: %s", address)
	return true
}

func (tp *TxParser) listenForTransactions(address string, txChannel chan models.Transaction) {
	defer close(txChannel)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(address)},
	}

	var logs []types.Log
	err := tp.client.CallContext(context.Background(), &logs, "eth_getLogs", query)
	if err != nil {
		log.Printf("Failed to fetch logs for address %s: %v", address, err)
		return
	}

	for _, logEntry := range logs {
		transaction, err := tp.getTransaction(logEntry.TxHash.Hex())
		if err != nil {
			log.Printf("Failed to get transaction for hash %s: %v", logEntry.TxHash.Hex(), err)
			continue
		}
		transaction.BlockNumber = strconv.FormatUint(logEntry.BlockNumber, 10)
		txChannel <- *transaction
	}
}

func (tp *TxParser) getTransaction(txHash string) (*models.Transaction, error) {
	tx := new(models.Transaction)
	err := tp.client.CallContext(context.Background(), &tx, "eth_getTransactionByHash", txHash)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch transaction details: %w", err)
	}

	return tx, nil
}
