package handler

import (
	"context"
	"gtihub.com/PowderDev/blockchain-indexer/store"

	"sync"

	"github.com/avast/retry-go/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/rs/zerolog/log"
	"gtihub.com/PowderDev/blockchain-indexer/addresstype"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type Options struct {
	Store     store.RedisConn
	EthClient *ethclient.Client
	ChainId   uint64
	Addresses map[common.Address]string
}

type Handler struct {
	store     store.RedisConn
	ethClient *ethclient.Client
	chainId   uint64
	addresses map[common.Address]string

	blockHeaders     map[common.Hash]*types.Header
	blockHeadersLock sync.Mutex
}

type EventMessage struct {
	ID        uint64                 `json:"id"`
	Type      string                 `json:"type"`
	Name      string                 `json:"name"`
	Timestamp uint64                 `json:"timestamp"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

type ContractHandler interface {
	Handle(ctx context.Context) error
	Channel() chan<- types.Log
}

func New(opts Options) *Handler {
	h := &Handler{
		store:        opts.Store,
		ethClient:    opts.EthClient,
		chainId:      opts.ChainId,
		addresses:    opts.Addresses,
		blockHeaders: map[common.Hash]*types.Header{},
	}
	return h
}

func (h *Handler) Handle(ctx context.Context, logs []types.Log) error {
	handlers := map[string]ContractHandler{
		addresstype.MSCEngine: NewMSCEngineHandler(h),
		addresstype.MakNFT:    NewMakNFTHandler(h),
	}

	g, gctx := errgroup.WithContext(ctx)

	g.Go(
		func() error {
			defer func() {
				for key := range handlers {
					close(handlers[key].Channel())
				}
			}()

			handled := map[LogId]bool{}

			for _, l := range logs {
				addressType, ok := h.addresses[l.Address]
				if !ok {
					return errors.Errorf("failed to get address type of %v", l.Address)
				}

				if l.Removed {
					log.Warn().Msg("log was removed")
					continue
				}

				logId := NewLogId(l)
				if _, ok := handled[logId]; ok {
					// already handled
					continue
				}

				handled[logId] = true

				select {
				case <-gctx.Done():
					return gctx.Err()
				case handlers[addressType].Channel() <- l:
				}
			}

			return nil
		},
	)

	for addressType := range handlers {
		handler := handlers[addressType]
		g.Go(
			func() error {
				return handler.Handle(gctx)
			},
		)
	}

	return g.Wait()
}

func (h *Handler) getBlockHeader(ctx context.Context, hash common.Hash) (*types.Header, error) {
	h.blockHeadersLock.Lock()
	defer h.blockHeadersLock.Unlock()
	header, ok := h.blockHeaders[hash]
	if ok {
		return header, nil
	}

	if err := retry.Do(
		func() error {
			var err error
			header, err = h.ethClient.HeaderByHash(ctx, hash)
			return err
		}, retry.Context(ctx),
	); err != nil {
		return nil, err
	}

	h.blockHeaders[hash] = header
	return header, nil
}
