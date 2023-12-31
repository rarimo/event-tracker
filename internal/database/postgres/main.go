package postgres

import (
	"github.com/rarimo/event-tracker/internal/database"
	"gitlab.com/distributed_lab/kit/pgdb"
)

// serviceDb is a struct that implements database.Database interface
type serviceDb struct {
	db *pgdb.DB
}

// NewServiceDB returns a new database.Database
func NewServiceDB(db *pgdb.DB) database.Database {
	return &serviceDb{db: db}
}

func (s *serviceDb) KeyValueQueryer() database.KeyValueQueryer {
	return NewKeyValueQ(s.db)
}

func (s *serviceDb) ProcessedTxQ() database.ProcessedTxQ {
	return NewProcessedTxsQ(s.db)
}
