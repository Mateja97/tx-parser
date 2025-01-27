package main

import (
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"tx-parser/api"
	"tx-parser/config"
	"tx-parser/parser"
)

func main() {

	config.Init()
	client, err := rpc.Dial(config.EthAddress())
	if err != nil {
		log.Fatalf("failed to connect to ethereum client: %v", err)
	}
	parser.Init(
		parser.WithClient(client))

	api.Init(api.WithPort(config.ServerAddress()))
	if err := api.ServerRun(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
