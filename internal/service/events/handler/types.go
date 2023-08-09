package handler

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type (
	// HandleFunc is a function that handles an event by a log.
	HandleFunc func(ctx context.Context, log types.Log) error
	// HandlersMap is a map of event handlers.
	HandlersMap map[string]HandleFunc

	// EventHandler is an interface for event handlers.
	EventHandler interface {
		// Handle handles event by a log.
		Handle(ctx context.Context, log types.Log) error
		// HandledTopics returns a list of topics that are handled by the handler.
		HandledTopics() []common.Hash
	}
)
