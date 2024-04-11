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

type MSCHEngineHandler struct {
	*Handler
	ch chan types.Log
}

func NewMSCEngineHandler(h *Handler) *MSCHEngineHandler {
	return &MSCHEngineHandler{
		Handler: h,
		ch:      make(chan types.Log),
	}

}

func (h *MSCHEngineHandler) Channel() chan<- types.Log {
	return h.ch
}

func (h *MSCHEngineHandler) Handle(ctx context.Context) error {
	gm := NewGroupManager(ctx)

	for l := range h.ch {
		switch l.Topics[0] {
		case MSCEngineAbi.Events["CollateralDeposited"].ID:
			event := &abi_gen.MSCEngineCollateralDeposited{
				Raw: l,
			}
			if err := UnpackMSCEngineLog(event, "CollateralDeposited", l); err != nil {
				return err
			}
			g, gctx := gm.Get(event.Depositor, event.Token)
			g.Go(
				func() error {
					return h.handleCollateralDeposited(gctx, event)
				},
			)

		case MSCEngineAbi.Events["CollateralRedeemed"].ID:
			event := &abi_gen.MSCEngineCollateralRedeemed{
				Raw: l,
			}
			if err := UnpackMSCEngineLog(event, "CollateralRedeemed", l); err != nil {
				return err
			}
			g, gctx := gm.Get(event.Redeemer, event.Token)
			g.Go(
				func() error {
					return h.handleCollateralRedeemed(gctx, event)
				},
			)

		case MSCEngineAbi.Events["DebtMinted"].ID:
			event := &abi_gen.MSCEngineDebtMinted{
				Raw: l,
			}
			if err := UnpackMSCEngineLog(event, "DebtMinted", l); err != nil {
				return err
			}
			g, gctx := gm.Get(event.To)
			g.Go(
				func() error {
					return h.handleDebtMinted(gctx, event)
				},
			)

		case MSCEngineAbi.Events["DebtBurned"].ID:
			event := &abi_gen.MSCEngineDebtBurned{
				Raw: l,
			}
			if err := UnpackMSCEngineLog(event, "DebtBurned", l); err != nil {
				return err
			}
			g, gctx := gm.Get(event.From)
			g.Go(
				func() error {
					return h.handleDebtBurned(gctx, event)
				},
			)

		case MSCEngineAbi.Events["Liquidated"].ID:
			event := &abi_gen.MSCEngineLiquidated{
				Raw: l,
			}
			if err := UnpackMSCEngineLog(event, "Liquidated", l); err != nil {
				return err
			}
			g, gctx := gm.Get(event.User, event.CollateralAddress)
			g.Go(
				func() error {
					return h.handleLiquidated(gctx, event)
				},
			)
		}
	}

	if err := gm.Wait(); err != nil {
		return err
	}

	return nil
}

func (h *MSCHEngineHandler) handleCollateralDeposited(
	ctx context.Context,
	event *abi_gen.MSCEngineCollateralDeposited,
) error {
	start := time.Now()

	blockHeader, err := h.getBlockHeader(ctx, event.Raw.BlockHash)
	if err != nil {
		return err
	}

	eventMessage := EventMessage{
		ID:        rand.Uint64(),
		Type:      "event",
		Name:      "CollateralDeposited",
		Timestamp: blockHeader.Time,
		Data: map[string]interface{}{
			"depositor":         event.Depositor.Hex(),
			"collateralAddress": event.Token.Hex(),
			"amount":            event.Amount.String(),
		},
	}

	encodedMessage, err := json.Marshal(eventMessage)
	if err != nil {
		return err
	}

	rKey := "event:" + eventMessage.Name + ":" + event.Depositor.Hex() + ":" + event.Token.Hex()
	err = h.store.Set(ctx, rKey, encodedMessage, 0).Err()
	if err != nil {
		return err
	}

	err = h.store.Publish(ctx, "events", encodedMessage).Err()
	if err != nil {
		return err
	}

	log.Info().Msgf("CollateralDeposited - %v", time.Since(start))

	return nil
}

func (h *MSCHEngineHandler) handleCollateralRedeemed(
	ctx context.Context,
	event *abi_gen.MSCEngineCollateralRedeemed,
) error {
	start := time.Now()

	blockHeader, err := h.getBlockHeader(ctx, event.Raw.BlockHash)
	if err != nil {
		return err
	}

	eventMessage := EventMessage{
		ID:        rand.Uint64(),
		Type:      "event",
		Name:      "CollateralRedeemed",
		Timestamp: blockHeader.Time,
		Data: map[string]interface{}{
			"redeemer":          event.Redeemer.Hex(),
			"collateralAddress": event.Token.Hex(),
			"amount":            event.Amount.String(),
		},
	}

	encodedMessage, err := json.Marshal(eventMessage)
	if err != nil {
		return err
	}

	rKey := "event:" + eventMessage.Name + ":" + event.Redeemer.Hex() + ":" + event.Token.Hex()
	err = h.store.Set(ctx, rKey, encodedMessage, 0).Err()
	if err != nil {
		return err
	}

	err = h.store.Publish(ctx, "events", encodedMessage).Err()
	if err != nil {
		return err
	}

	log.Info().Msgf("CollateralRedeemed - %v", time.Since(start))

	return nil
}

func (h *MSCHEngineHandler) handleDebtMinted(ctx context.Context, event *abi_gen.MSCEngineDebtMinted) error {
	start := time.Now()

	blockHeader, err := h.getBlockHeader(ctx, event.Raw.BlockHash)
	if err != nil {
		return err
	}

	eventMessage := EventMessage{
		ID:        rand.Uint64(),
		Type:      "event",
		Name:      "DebtMinted",
		Timestamp: blockHeader.Time,
		Data: map[string]interface{}{
			"mintedTo": event.To.Hex(),
			"amount":   event.Amount.String(),
		},
	}

	encodedMessage, err := json.Marshal(eventMessage)
	if err != nil {
		return err
	}

	rKey := "event:" + eventMessage.Name + ":" + event.To.Hex()
	err = h.store.Set(ctx, rKey, encodedMessage, 0).Err()
	if err != nil {
		return err
	}

	err = h.store.Publish(ctx, "events", encodedMessage).Err()
	if err != nil {
		return err
	}

	log.Info().Msgf("DebtMinted - %v", time.Since(start))

	return nil
}

func (h *MSCHEngineHandler) handleDebtBurned(ctx context.Context, event *abi_gen.MSCEngineDebtBurned) error {
	start := time.Now()

	blockHeader, err := h.getBlockHeader(ctx, event.Raw.BlockHash)
	if err != nil {
		return err
	}

	eventMessage := EventMessage{
		ID:        rand.Uint64(),
		Type:      "event",
		Name:      "DebtBurned",
		Timestamp: blockHeader.Time,
		Data: map[string]interface{}{
			"burnedFrom": event.From.Hex(),
			"amount":     event.Amount.String(),
		},
	}

	encodedMessage, err := json.Marshal(eventMessage)
	if err != nil {
		return err
	}

	rKey := "event:" + eventMessage.Name + ":" + event.From.Hex()
	err = h.store.Set(ctx, rKey, encodedMessage, 0).Err()
	if err != nil {
		return err
	}

	err = h.store.Publish(ctx, "events", encodedMessage).Err()
	if err != nil {
		return err
	}

	log.Info().Msgf("DebtBurned - %v", time.Since(start))

	return nil
}

func (h *MSCHEngineHandler) handleLiquidated(ctx context.Context, event *abi_gen.MSCEngineLiquidated) error {
	start := time.Now()

	blockHeader, err := h.getBlockHeader(ctx, event.Raw.BlockHash)
	if err != nil {
		return err
	}

	eventMessage := EventMessage{
		ID:        rand.Uint64(),
		Type:      "event",
		Name:      "Liquidated",
		Timestamp: blockHeader.Time,
		Data: map[string]interface{}{
			"user":              event.User.Hex(),
			"collateralAddress": event.CollateralAddress.Hex(),
			"amount":            event.Amount.String(),
		},
	}

	encodedMessage, err := json.Marshal(eventMessage)
	if err != nil {
		return err
	}

	rKey := "event:" + eventMessage.Name + ":" + event.User.Hex() + ":" + event.CollateralAddress.Hex()
	err = h.store.Set(ctx, rKey, encodedMessage, 0).Err()
	if err != nil {
		return err
	}

	err = h.store.Publish(ctx, "events", encodedMessage).Err()
	if err != nil {
		return err
	}

	log.Info().Msgf("Liquidated - %v", time.Since(start))

	return nil
}
