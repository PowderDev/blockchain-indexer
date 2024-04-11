package main

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	RedisUrl             string `envconfig:"REDIS_URL" required:"true"`
	RpcUrl               string `envconfig:"INDEXER_RPC_URL" required:"true"`
	Version              uint64 `envconfig:"INDEXER_VERSION" required:"true"`
	NFTAddress           string `envconfig:"INDEXER_NFT_ADDRESS" required:"true"`
	MSCEngineAddress     string `envconfig:"INDEXER_MSC_ENGINE_ADDRESS" required:"true"`
	StartBlock           uint64 `envconfig:"INDEXER_START_BLOCK" required:"true"`
	BlockRange           uint64 `envconfig:"INDEXER_BLOCK_RANGE" required:"true"`
	BlockPollingInterval uint64 `envconfig:"INDEXER_BLOCK_POLLING_INTERVAL" required:"true"`
	LogLevel             string `envconfig:"INDEXER_LOG_LEVEL" default:"trace"`
	ConfirmationsNeeded  uint64 `envconfig:"INDEXER_CONFIRMATIONS_NEEDED" required:"true"`
}

func LoadConfig() (Config, error) {
	var c Config
	return c, envconfig.Process("", &c)
}
