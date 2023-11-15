package handler

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rarimo/event-tracker/internal/contracts"
	"github.com/rarimo/event-tracker/internal/service/galxe"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"strings"
)

const (
	IdentityProved = "IdentityProved"
)

// eventHandler is an implementation of EventHandler.
type eventHandler struct {
	logger *logan.Entry

	contractAbi     abi.ABI
	supportedEvents []common.Hash

	galxeSubmitter galxe.Submitter
	skipSubmitter  bool
}

// MustNewEventHandler creates a new instance of EventHandler.
// It panics if an error occurs.
func MustNewEventHandler(logger *logan.Entry, submitter galxe.Submitter) EventHandler {
	contractAbi, err := abi.JSON(strings.NewReader(contracts.DemoVerifierMetaData.ABI))
	if err != nil {
		panic(err)
	}

	return &eventHandler{
		galxeSubmitter:  submitter,
		contractAbi:     contractAbi,
		supportedEvents: []common.Hash{contractAbi.Events[IdentityProved].ID},
		logger:          logger.WithField("component", "event-handler"),
	}
}

// Handle handles event by a log.
func (h *eventHandler) Handle(ctx context.Context, log types.Log) error {
	// First topic must be a hashed signature of the event
	topic := log.Topics[0]

	event, err := h.contractAbi.EventByID(topic)
	if err != nil {
		return errors.Wrap(err, "failed to get event by topic", logan.F{"topic": topic.Hex()})
	}

	handler, exists := h.Handlers()[event.Name]
	if !exists {
		return errors.Wrap(err, "event handler not found", logan.F{"event": event.Name})
	}

	return handler(ctx, log)
}

// Handlers returns a map of event handlers.
func (h *eventHandler) Handlers() HandlersMap {
	return HandlersMap{
		IdentityProved: h.HandleIdentityProvedEvent,
	}
}

// HandledTopics returns a list of topics that are handled by the handler.
func (h *eventHandler) HandledTopics() []common.Hash {
	return h.supportedEvents
}
