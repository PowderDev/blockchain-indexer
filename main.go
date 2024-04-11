package main

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/redis/go-redis/v9"
	"math/big"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"gtihub.com/PowderDev/blockchain-indexer/addresstype"
	"gtihub.com/PowderDev/blockchain-indexer/handler"

	"gtihub.com/PowderDev/blockchain-indexer/store"

	_ "github.com/joho/godotenv/autoload"
)

var (
	cfg Config
)

func init() {
	var err error

	cfg, err = LoadConfig()
	if err != nil {
		panic(err)
	}

	logLevel, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	zerolog.SetGlobalLevel(logLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

func main() {
	if err := run(); err != nil {
		log.Fatal().Stack().Err(err).Msg("")
	}
}

func run() error {
	ctx := context.Background()

	r := redis.NewClient(
		&redis.Options{
			Addr:     cfg.RedisUrl,
			Password: "",
			DB:       0,
		},
	)

	ethClient, err := ethclient.DialContext(ctx, cfg.RpcUrl)
	if err != nil {
		return errors.WithStack(err)
	}
	defer ethClient.Close()

	var chainId uint64
	if res, err := ethClient.ChainID(ctx); err != nil {
		return errors.WithStack(err)
	} else {
		chainId = res.Uint64()
	}

	fromBlock, err := store.GetCurrentBlock(ctx, r, chainId, cfg.Version, cfg.StartBlock)
	if err != nil {
		return errors.WithStack(err)
	}

	addresses := map[common.Address]string{
		common.HexToAddress(cfg.NFTAddress):       addresstype.MakNFT,
		common.HexToAddress(cfg.MSCEngineAddress): addresstype.MSCEngine,
	}

	h := handler.New(
		handler.Options{
			Store:     r,
			EthClient: ethClient,
			ChainId:   chainId,
			Addresses: addresses,
		},
	)

	for {
		latestBlockHeader, err := ethClient.HeaderByNumber(ctx, nil)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get latest block")
			sleepForBlockInterval()
			continue
		}

		latestBlock := latestBlockHeader.Number.Uint64()
		latestBlock -= cfg.ConfirmationsNeeded

		blockRange := cfg.BlockRange

		if fromBlock > latestBlock {
			log.Info().Msg("No new blocks")
			handleNoNewBlocks(r, ctx)
			sleepForBlockInterval()
			continue
		}

		for fromBlock <= latestBlock {
			fromBlock, blockRange = handleBlock(
				r, ctx, ethClient, chainId, addresses, fromBlock, latestBlock, blockRange, h,
			)
		}
	}
}

func handleBlock(
	r *redis.Client,
	globalCtx context.Context,
	ethClient *ethclient.Client,
	chainId uint64,
	addresses map[common.Address]string,
	fromBlock uint64,
	latestBlock uint64,
	blockRange uint64,
	h *handler.Handler,
) (uint64, uint64) {
	toBlock := fromBlock + blockRange
	if toBlock > latestBlock {
		toBlock = latestBlock
	}

	log.Info().Msgf("Scanning blocks [%v, %v]", fromBlock, toBlock)

	filterQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
	}

	filterQuery.Addresses = make([]common.Address, 0, len(addresses))
	for address := range addresses {
		filterQuery.Addresses = append(filterQuery.Addresses, address)
	}

	logs, err := ethClient.FilterLogs(globalCtx, filterQuery)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get logs")
		// reduce block range to avoid "query returned more than 10000 results" error
		blockRange = blockRange / 2
		sleepForBlockInterval()
		return fromBlock, blockRange
	}

	log.Info().Msgf("Logs: %v, Block: %v", len(logs), fromBlock)

	fromBlock, blockRange, err = handleLogs(r, globalCtx, chainId, logs, toBlock, h)

	if err != nil {
		log.Error().Stack().Err(err).Msg("")
	}

	return fromBlock, blockRange
}

func handleLogs(
	r *redis.Client,
	globalCtx context.Context,
	chainId uint64,
	logs []types.Log,
	toBlock uint64,
	h *handler.Handler,
) (uint64, uint64, error) {
	ctx, cancel := context.WithCancel(globalCtx)
	defer cancel()

	tx := r.Pipeline()

	newFromBlock := toBlock + 1

	if err := h.Handle(ctx, logs); err != nil {
		return newFromBlock, cfg.BlockRange, errors.WithStack(err)
	}

	if err := store.UpdateCurrentBlock(ctx, tx, chainId, cfg.Version, newFromBlock); err != nil {
		return newFromBlock, cfg.BlockRange, errors.WithStack(err)
	}

	if len(logs) == 0 {
		handleNoNewLogs(r, ctx)
	}

	if _, err := tx.Exec(ctx); err != nil {
		return 0, 0, errors.WithStack(err)
	}

	return newFromBlock, cfg.BlockRange, nil
}

func handleNoNewLogs(r store.RedisConn, ctx context.Context) {
	eventMessage := handler.EventMessage{
		ID:        rand.Uint64(),
		Type:      "no-new-logs",
		Timestamp: uint64(time.Now().Unix()),
		Name:      "New block but no new relevant logs",
	}

	encodedMessage, _ := json.Marshal(eventMessage)
	r.Publish(ctx, "events", encodedMessage)
}

func handleNoNewBlocks(r store.RedisConn, ctx context.Context) {
	eventMessage := handler.EventMessage{
		ID:        rand.Uint64(),
		Type:      "no-new-blocks",
		Timestamp: uint64(time.Now().Unix()),
		Name:      "No new blocks",
	}

	encodedMessage, _ := json.Marshal(eventMessage)
	r.Publish(ctx, "events", encodedMessage)
}

func sleepForBlockInterval() {
	time.Sleep(time.Millisecond * time.Duration(cfg.BlockPollingInterval))
}
