package ctxgrp

import (
	"context"
	"sync"
)

// WaitGroup interface for built-in WaitGroup
type WaitGroup interface {
	Add(delta int)
	Done()
	Wait()
}

// Group provides context for a group of goroutines,
// a parent context and a WaitGroup for waiting for children to finish;
// is used for membership management & supervision (cancellation)
type Group func() (context.Context, WaitGroup)

// New if ctx is not provided, context.Background() will be used
func New(ctx ...context.Context) Group {
	var x context.Context
	if len(ctx) > 0 && ctx[0] != nil {
		x = ctx[0]
	} else {
		x = context.Background()
	}
	return (&group{ctx: x, wg: &sync.WaitGroup{}}).elements
}

type group struct {
	ctx context.Context
	wg  *sync.WaitGroup
}

func (g *group) elements() (context.Context, WaitGroup) {
	return g.ctx, g.wg
}
