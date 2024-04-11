package handler

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gtihub.com/PowderDev/blockchain-indexer/abi_gen"
)

var (
	MakNFTAbi    abi.ABI
	MSCEngineAbi abi.ABI
)

func init() {
	var err error

	MSCEngineAbi, err = abi.JSON(strings.NewReader(abi_gen.MSCEngineMetaData.ABI))
	if err != nil {
		panic(err)
	}

	MakNFTAbi, err = abi.JSON(strings.NewReader(abi_gen.MakNFTMetaData.ABI))
	if err != nil {
		panic(err)
	}

}

func UnpackMSCEngineLog(out interface{}, event string, l types.Log) error {
	return bind.NewBoundContract(common.Address{}, MSCEngineAbi, nil, nil, nil).UnpackLog(out, event, l)
}

func UnpackMakNFTLog(out interface{}, event string, l types.Log) error {
	return bind.NewBoundContract(common.Address{}, MakNFTAbi, nil, nil, nil).UnpackLog(out, event, l)
}
