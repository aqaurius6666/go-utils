package cockroach

import (
	"context"

	"github.com/aqaurius6666/go-utils/go1.18/database"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

var (
	_ database.CommonRepository = (*CDBRepository)(nil)
)

type DBInterfaces []interface{}

type CDBRepository struct {
	Db         *gorm.DB
	Logger     *logrus.Logger
	Context    context.Context
	Interfaces DBInterfaces
}

func InitCDBRepository(ctx context.Context, logger *logrus.Logger, db *gorm.DB) CDBRepository {
	return CDBRepository{
		Db:         db,
		Logger:     logger,
		Context:    ctx,
		Interfaces: nil,
	}
}

func (c *CDBRepository) RawSQL(sql string, args ...interface{}) error {
	return c.Db.Exec(sql, args...).Error
}

func (c *CDBRepository) SetInterfaces(itf ...interface{}) {
	c.Interfaces = itf
}
func (c *CDBRepository) Close() error {
	d, err := c.Db.DB()
	if err != nil {
		return err
	}
	return d.Close()
}

func (c *CDBRepository) Migrate() error {
	if c.Interfaces == nil {
		return xerrors.New("empty interfaces")
	}
	return c.Db.AutoMigrate(c.Interfaces...)
}

func (c *CDBRepository) Drop() error {
	if c.Interfaces == nil {
		return xerrors.New("empty interfaces")
	}
	return c.Db.Migrator().DropTable(c.Interfaces...)
}
