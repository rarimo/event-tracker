package postgres

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/rarimo/event-tracker/internal/database"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const (
	keyValueTable = "key_value"

	keyColumn   = "key"
	valueColumn = "value"
)

var keyValueSelect = squirrel.Select("*").From(keyValueTable)

type keyValueQ struct {
	db *pgdb.DB
}

func NewKeyValueQ(db *pgdb.DB) database.KeyValueQueryer {
	return &keyValueQ{
		db: db,
	}
}

func (q *keyValueQ) Upsert(kv database.KeyValue) error {
	query := squirrel.Insert(keyValueTable).
		SetMap(structs.Map(kv)).
		Suffix("ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value")

	return q.db.Exec(query)
}

func (q *keyValueQ) New() database.KeyValueQueryer {
	return NewKeyValueQ(q.db.Clone())
}

func (q *keyValueQ) Get(key string) (*database.KeyValue, error) {
	return q.get(key, false)
}

func (q *keyValueQ) LockingGet(key string) (*database.KeyValue, error) {
	return q.get(key, true)
}

func (q *keyValueQ) get(key string, forUpdate bool) (*database.KeyValue, error) {
	statement := keyValueSelect.Where(squirrel.Eq{keyColumn: key})
	if forUpdate {
		statement = statement.Suffix("FOR UPDATE")
	}

	var value database.KeyValue
	err := q.db.Get(&value, statement)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &value, err
}
