package parser

import "github.com/ethereum/go-ethereum/rpc"

func WithClient(client *rpc.Client) func(*TxParser) {
	return func(parser *TxParser) {
		parser.client = client
	}
}
