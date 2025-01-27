package parser

import (
	"context"
	"log"
	"math/big"
)

func (tp *TxParser) GetCurrentBlock() int {
	var blockNumber string
	err := tp.client.CallContext(context.Background(), &blockNumber, "eth_blockNumber")
	if err != nil {
		log.Printf("Failed to fetch current block number: %v", err)
		return 0
	}

	block, ok := new(big.Int).SetString(blockNumber[2:], 16)
	if !ok {
		log.Printf("Failed to convert block number: %v", blockNumber)
		return 0
	}
	return int(block.Int64())
}
