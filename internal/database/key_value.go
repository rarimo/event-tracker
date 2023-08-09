package database

// KeyValue represents a key-value pair
type KeyValue struct {
	Key   string `db:"key" structs:"key"`
	Value string `db:"value" structs:"value"`
}

// KeyValueQueryer is a query interface for KeyValue
//
//go:generate mockery --case=underscore --name=KeyValueQueryer
type KeyValueQueryer interface {
	// New creates a new KeyValueQueryer
	New() KeyValueQueryer
	// Get reads row and does not lock the row
	Get(key string) (*KeyValue, error)
	// Upsert updates value if there is one, insert if no
	Upsert(KeyValue) error
	// LockingGet reads row and locks the row for reading and updating
	// until the end of the current transaction
	LockingGet(key string) (*KeyValue, error)
}
