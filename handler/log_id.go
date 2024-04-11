package handler

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type LogId struct {
	blockNumber uint64
	txHash      common.Hash
	index       uint
}

func NewLogId(l types.Log) LogId {
	return LogId{
		blockNumber: l.BlockNumber,
		txHash:      l.TxHash,
		index:       l.Index,
	}
}
