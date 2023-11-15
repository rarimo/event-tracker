package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	data "github.com/rarimo/event-tracker/internal/database"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const processedTxsTableName = "processed_txs"

var defaultSelector = squirrel.Select("*").From(processedTxsTableName)

func NewProcessedTxsQ(db *pgdb.DB) data.ProcessedTxQ {
	return &processedTxsQ{
		db:       db,
		selector: defaultSelector,
		updater:  squirrel.Update(processedTxsTableName),
	}
}

type processedTxsQ struct {
	db       *pgdb.DB
	selector squirrel.SelectBuilder
	updater  squirrel.UpdateBuilder
}

func (q *processedTxsQ) Get() (*data.ProcessedTx, error) {
	var result data.ProcessedTx

	stmt := q.selector
	// dropping selector to default
	q.selector = defaultSelector

	err := q.db.Get(&result, stmt)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (q *processedTxsQ) OrderBy(column, order string) data.ProcessedTxQ {
	q.selector = q.selector.OrderBy(fmt.Sprintf("%s %s", column, order))
	return q
}

func (q *processedTxsQ) Limit(limit uint64) data.ProcessedTxQ {
	q.selector = q.selector.Limit(limit)
	return q
}

func (q *processedTxsQ) FilterByHash(hash string) data.ProcessedTxQ {
	q.selector = q.selector.Where(squirrel.Eq{"hash": hash})
	return q
}

func (q *processedTxsQ) Insert(value data.ProcessedTx) error {
	clauses := structs.Map(value)
	var result uint64
	stmt := squirrel.Insert(processedTxsTableName).SetMap(clauses).Suffix("returning id")
	err := q.db.Get(&result, stmt)
	return err
}

func (q *processedTxsQ) TxExists(hash string) bool {
	var result data.ProcessedTx
	stmt := q.selector.Where(squirrel.Eq{"hash": hash})
	_ = q.db.Get(&result, stmt)
	return result != data.ProcessedTx{}
}
