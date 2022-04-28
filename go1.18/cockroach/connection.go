package cockroach

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	GORM_DEFAULT_CONFIG = &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
	}
)

func NewCDBConnection(dsn string, cfg *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), cfg)
	if err != nil {
		return nil, err
	}
	return db, nil
}
