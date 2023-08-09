package handler

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/rarimo/identity/event-tracker/internal/contracts"
)

// HandleIdentityProvedEvent handles IdentityProved event.
func (h *eventHandler) HandleIdentityProvedEvent(_ context.Context, log types.Log) error {
	var event contracts.DemoVerifierIdentityProved
	if err := h.contractAbi.UnpackIntoInterface(&event, IdentityProved, log.Data); err != nil {
		return errors.Wrap(err, "failed to unpack event", logan.F{
			"event": IdentityProved,
		})
	}

	event.IdentityId = log.Topics[1].Big() // param is indexed and located in the log.Topics[1]

	fields := logan.F{
		"event":       IdentityProved,
		"tx":          log.TxHash.Hex(),
		"identity-id": event.IdentityId.String(),
		"sender":      event.SenderAddr.Hex(),
	}

	h.logger.WithFields(fields).Debug("Processing identity proved event...")

	eligible, err := h.galxeSubmitter.SubmitParticipant(event.SenderAddr)
	if err != nil {
		return errors.Wrap(err, "failed to submit participant to the Galxe API", fields)
	}

	if !eligible {
		h.logger.WithFields(fields).Info("Participant was not submitted to the Galxe API")
		return nil
	}

	h.logger.WithFields(fields).Info("Participant was successfully submitted to the Galxe API")

	return nil
}
