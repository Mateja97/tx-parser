package config

import (
	"os"
)

type Config struct {
	serverPort string
	ethAddress string
}

var c *Config

func Init() {
	c = new(Config)

	c.serverPort = os.Getenv("SERVER_PORT")
	c.ethAddress = os.Getenv("ETH_ADDRESS")

}

func ServerAddress() string {
	return c.serverPort
}

func EthAddress() string {
	return c.ethAddress
}
