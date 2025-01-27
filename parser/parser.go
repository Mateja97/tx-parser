package parser

import (
	"github.com/ethereum/go-ethereum/rpc"
	"sync"
	"tx-parser/models"
)

var _ Parser = &TxParser{}

type Parser interface {
	GetCurrentBlock() int
	Subscribe(address string) bool
	GetTransactions(address string) []models.Transaction
}

type TxParser struct {
	client         *rpc.Client
	subscribedAddr map[string]chan models.Transaction
	mu             sync.RWMutex
}

var tp *TxParser

func Init(options ...func(*TxParser)) {
	tp = new(TxParser)
	for _, o := range options {
		o(tp)
	}
	tp.subscribedAddr = make(map[string]chan models.Transaction)
}
func Subscribe(address string) bool {
	return tp.Subscribe(address)
}

func GetCurrentBlock() int {
	return tp.GetCurrentBlock()
}

func GetTransactions(address string) []models.Transaction {
	return tp.GetTransactions(address)
}
