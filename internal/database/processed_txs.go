package database

type ProcessedTxQ interface {
	Get() (*ProcessedTx, error)
	OrderBy(column, order string) ProcessedTxQ
	Limit(amt uint64) ProcessedTxQ
	FilterByHash(hash string) ProcessedTxQ
	Insert(tx ProcessedTx) error
	TxExists(hash string) bool
}

type ProcessedTx struct {
	ID    string `db:"id" structs:"-"`
	Block uint64 `db:"block" structs:"block"`
	Hash  string `db:"hash" structs:"hash"`
}
