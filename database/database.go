package database

type CommonRepository interface {
	Close() error
	Migrate() error
	Drop() error
	RawSQL(string, ...interface{}) error
}
