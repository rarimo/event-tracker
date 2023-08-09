package tracker

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
	"gitlab.com/rarimo/identity/event-tracker/internal/config"
	"gitlab.com/rarimo/identity/event-tracker/internal/database"
	"gitlab.com/rarimo/identity/event-tracker/internal/service/events/handler"
	"math/big"
	"strconv"
)

// eventTracker is a struct that implements EventTracker
type eventTracker struct {
	logger       *logan.Entry
	settings     config.EventTrackerSettings
	db           database.Database
	eventHandler handler.EventHandler
}

// NewEventTracker returns new instance of EventTracker
func NewEventTracker(
	settings config.EventTrackerSettings,
	logger *logan.Entry,
	db database.Database,
	eventHandler handler.EventHandler,
) EventTracker {
	return &eventTracker{
		logger:       logger.WithField("service", settings.Name),
		settings:     settings,
		db:           db,
		eventHandler: eventHandler,
	}
}

// Run starts the event tracker service.
// It is a blocking function as it uses a running.WithBackOff.
func (t *eventTracker) Run(ctx context.Context) {
	t.logger.Info(fmt.Sprintf("started %s runner", t.settings.Name))

	running.WithBackOff(
		ctx,
		t.logger,
		t.settings.Name,
		t.ProcessInterval,
		t.settings.NormalPeriod,
		t.settings.MinAbnormalPeriod,
		t.settings.MaxAbnormalPeriod,
	)
}

// ProcessInterval tracks and processes events for a given interval of blocks.
// Calculates the range of blocks to process based on the last processed block number and iteration size.
func (t *eventTracker) ProcessInterval(ctx context.Context) error {
	fromBlock, toBlock, err := t.getBlockInterval()
	if err != nil {
		return errors.Wrap(err, "failed to get block interval")
	}

	t.logger.Debug(fmt.Sprintf("processing interval from %d to %d", fromBlock, toBlock))

	logs, err := t.settings.RPC.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(fromBlock),
		ToBlock:   new(big.Int).SetUint64(toBlock),
		Addresses: []common.Address{t.settings.ContractAddress},
		Topics:    [][]common.Hash{t.eventHandler.HandledTopics()},
	})

	t.logger.Debug(fmt.Sprintf("found %d logs", len(logs)))

	for _, log := range logs {
		if err := t.ProcessLog(ctx, log); err != nil {
			return errors.Wrap(err, "failed to process log")
		}
	}

	if err := t.db.KeyValueQueryer().Upsert(database.KeyValue{
		Key:   t.settings.Name,
		Value: strconv.FormatInt(int64(toBlock), 10),
	}); err != nil {
		return errors.Wrap(err, "failed to save last processed")
	}

	return nil
}

// ProcessLog processes a single log and saves it to the database if it was not processed before.
func (t *eventTracker) ProcessLog(ctx context.Context, log types.Log) error {
	uidHash := getLogUniqueIdentifier(log)

	tx, err := t.db.ProcessedTxQ().FilterByHash(uidHash).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get processed tx")
	}
	if tx != nil {
		t.logger.Debug(fmt.Sprintf("tx %s already processed, omitting", uidHash))
		return nil
	}

	if err = t.eventHandler.Handle(ctx, log); err != nil {
		return errors.Wrap(err, "failed to handle event")
	}

	if err = t.db.ProcessedTxQ().Insert(database.ProcessedTx{
		Hash:  uidHash,
		Block: log.BlockNumber,
	}); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to insert processed tx %s", uidHash))
	}

	return nil
}

// getBlockInterval returns the range of blocks to process.
func (t *eventTracker) getBlockInterval() (fromBlock, toBlock uint64, err error) {
	lastBlock, err := t.getLastProcessedBlock()
	if err != nil {
		return fromBlock, toBlock, errors.Wrap(err, "failed to get last processed block")
	}

	nextBlock, err := t.getNextBlock(lastBlock)
	if err != nil {
		return fromBlock, toBlock, errors.Wrap(err, "failed to get next block")
	}

	return lastBlock, nextBlock, nil
}

// getNextBlock returns the next block number to process.
func (t *eventTracker) getNextBlock(lastProcessedBlock uint64) (uint64, error) {
	// Retrieving the last blockchain block number
	lastBlock, err := t.settings.RPC.BlockNumber(context.Background())
	if err != nil {
		return 0, errors.Wrap(err, "failed to get the last block in the blockchain")
	}

	nextBlock := lastProcessedBlock + t.settings.IterationSize
	if nextBlock > lastBlock {
		nextBlock = lastBlock
	}

	return nextBlock, nil
}

// getLastProcessedBlock returns the last processed block number.
func (t *eventTracker) getLastProcessedBlock() (uint64, error) {
	cursor, err := t.db.KeyValueQueryer().LockingGet(t.settings.Name)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get current cursor value", logan.F{"key": t.settings.Name})
	}

	// If cursor is nil, then we have not processed any blocks yet.
	// In this case, we should return the start block.
	if cursor == nil {
		return t.settings.StartBlock, nil
	}

	lastBlock, err := strconv.ParseInt(cursor.Value, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse cursor", logan.F{"kv_cursor": cursor.Value})
	}

	// should never happen, but just in case
	if lastBlock < 0 {
		return 0, errors.From(errors.New("block number cannot be negative"), logan.F{
			"block_number": lastBlock,
		})
	}

	return uint64(lastBlock), nil
}

// getLogUniqueIdentifier returns a unique identifier for a given log.
func getLogUniqueIdentifier(log types.Log) string {
	return fmt.Sprintf("%s-%s", log.TxHash.String(), strconv.Itoa(int(log.Index)))
}
