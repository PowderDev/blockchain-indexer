package handler

import (
	"context"
	"encoding/json"
	"math/rand/v2"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
	"gtihub.com/PowderDev/blockchain-indexer/abi_gen"
)

type MakNFTHandler struct {
	*Handler
	ch chan types.Log
}

func NewMakNFTHandler(h *Handler) *MakNFTHandler {
	return &MakNFTHandler{
		Handler: h,
		ch:      make(chan types.Log),
	}
}

func (h *MakNFTHandler) Channel() chan<- types.Log {
	return h.ch
}

func (h *MakNFTHandler) Handle(ctx context.Context) error {
	gm := NewGroupManager(ctx)

	for l := range h.ch {
		switch l.Topics[0] {
		case MakNFTAbi.Events["NFTRequested"].ID:
			event := &abi_gen.MakNFTNFTRequested{
				Raw: l,
			}
			if err := UnpackMakNFTLog(event, "NFTRequested", l); err != nil {
				return err
			}

			g, gctx := gm.Get(event.Requester, event.RequestId)
			g.Go(
				func() error {
					return h.handleNFTRequested(gctx, event)
				},
			)

		case MakNFTAbi.Events["NFTMinted"].ID:
			event := &abi_gen.MakNFTNFTMinted{
				Raw: l,
			}
			if err := UnpackMakNFTLog(event, "NFTMinted", l); err != nil {
				return err
			}

			g, gctx := gm.Get(event.Minter, event.TokenId, event.Uri)
			g.Go(
				func() error {
					return h.handleNFTMinted(gctx, event)
				},
			)
		}
	}

	if err := gm.Wait(); err != nil {
		return err
	}

	return nil
}

func (h *MakNFTHandler) handleNFTRequested(ctx context.Context, event *abi_gen.MakNFTNFTRequested) error {
	start := time.Now()

	blockHeader, err := h.getBlockHeader(ctx, event.Raw.BlockHash)
	if err != nil {
		return err
	}

	eventMessage := EventMessage{
		ID:        rand.Uint64(),
		Type:      "event",
		Name:      "NFTRequested",
		Timestamp: blockHeader.Time,
		Data: map[string]interface{}{
			"requester": event.Requester.Hex(),
			"requestId": event.RequestId.String(),
		},
	}

	encodedMessage, err := json.Marshal(eventMessage)
	if err != nil {
		return err
	}

	rKey := "event:" + eventMessage.Name + ":" + event.Requester.Hex() + ":" + event.RequestId.String()
	err = h.store.Set(ctx, rKey, encodedMessage, 0).Err()
	if err != nil {
		return err
	}

	err = h.store.Publish(ctx, "events", encodedMessage).Err()
	if err != nil {
		return err
	}

	log.Info().Msgf("NFTRequested - %v", time.Since(start))

	return nil
}

func (h *MakNFTHandler) handleNFTMinted(ctx context.Context, event *abi_gen.MakNFTNFTMinted) error {
	start := time.Now()

	blockHeader, err := h.getBlockHeader(ctx, event.Raw.BlockHash)
	if err != nil {
		return err
	}

	eventMessage := EventMessage{
		ID:        rand.Uint64(),
		Type:      "event",
		Name:      "NFTMinted",
		Timestamp: blockHeader.Time,
		Data: map[string]interface{}{
			"minter":  event.Minter.Hex(),
			"tokenId": event.TokenId.String(),
			"uri":     event.Uri,
		},
	}

	encodedMessage, err := json.Marshal(eventMessage)
	if err != nil {
		return err
	}

	rKey := "event:" + eventMessage.Name + ":" + event.Minter.Hex() + ":" + event.TokenId.String()
	err = h.store.Set(ctx, rKey, encodedMessage, 0).Err()
	if err != nil {
		return err
	}

	err = h.store.Publish(ctx, "events", encodedMessage).Err()
	if err != nil {
		return err
	}

	log.Info().Msgf("NFTMinted - %v", time.Since(start))

	return nil
}
