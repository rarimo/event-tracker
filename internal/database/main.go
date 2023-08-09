package database

// Database is an interface for service database to access and query different tables.
type Database interface {
	// KeyValueQueryer returns a KeyValueQueryer interface for key_value table.
	KeyValueQueryer() KeyValueQueryer
	// ProcessedTxQ returns a ProcessedTxQ interface for processed_txs table.
	ProcessedTxQ() ProcessedTxQ
}
