package handler

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

type Group struct {
	g    *errgroup.Group
	gctx context.Context
}

// GroupManager is useful for creating a goroutine (errgroup with limit=1) for each token.
// This allows us to use concurrent handlers and avoid race conditions in some of them that depend on the previous state.
type GroupManager struct {
	groups map[string]*Group
	ctx    context.Context
}

func NewGroupManager(ctx context.Context) *GroupManager {
	return &GroupManager{
		groups: map[string]*Group{},
		ctx:    ctx,
	}
}

func (gm *GroupManager) Get(keyParts ...interface{}) (*errgroup.Group, context.Context) {
	key := fmt.Sprintf("%v", keyParts)
	if _, ok := gm.groups[key]; !ok {
		g, gctx := errgroup.WithContext(gm.ctx)
		g.SetLimit(1)
		gm.groups[key] = &Group{
			g:    g,
			gctx: gctx,
		}
	}
	return gm.groups[key].g, gm.groups[key].gctx
}

func (gm *GroupManager) Wait() error {
	for key := range gm.groups {
		if err := gm.groups[key].g.Wait(); err != nil {
			return err
		}
	}
	return nil
}
