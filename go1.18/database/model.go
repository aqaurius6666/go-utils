package database

import (
	"github.com/aqaurius6666/go-utils/utils"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	db *gorm.DB
)

type ApplyFunc func(*gorm.DB) *gorm.DB

type Applier interface {
	Apply(*gorm.DB) *gorm.DB
}

func ConnectDB(dsn string) error {
	inst, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return err
	}
	db = inst
	return nil
}
func Migrate(list ...interface{}) error {
	return db.AutoMigrate(list...)
}

type BaseEntityMustImplStr interface {
	StrBaseEntityMustImpl()
}
type ApplySearcher interface {
	Apply(*gorm.DB) *gorm.DB
}
type BaseEntityMustImplPtr interface {
	PtrBaseEntityMustImpl()
}
type EntityPtr interface {
	BaseEntitySetterGetter
	TableNamer
	BaseEntityMustImplPtr
}
type TableNamer interface {
	TableName() string
}
type BaseEntitySetterGetter interface {
	GetBaseEntity() *BaseEntity
	SetBaseEntity(*BaseEntity)
}

type RepositoryInterface[T EntityPtr] interface {
	SelectOne(Applier) (T, error)
	List(Applier) ([]T, error)
	CreateOne(T) (T, error)
}
type BaseEntity struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id" json:"id"`
	UpdatedAt int64     `gorm:"autoUpdateTime:milli" json:"updatedAt"`
	CreatedAt int64     `gorm:"autoCreateTime:milli" json:"createdAt"`
	DeletedAt int64     `gorm:"autoDeleteTime:milli" json:"deletedAt"`
}

func (s *BaseEntity) PtrBaseEntityMustImpl() {}
func (s BaseEntity) StrBaseEntityMustImpl()  {}

type BaseSearch struct {
	Skip      int
	Limit     int
	OrderBy   string
	OrderType string
	Fields    []string
}

type BaseSearchEntity[T EntityPtr] struct {
	Entity T
	BaseSearch
	applyFuncs []ApplyFunc
}

func (s *BaseSearchEntity[T]) Apply(db *gorm.DB) *gorm.DB {
	for _, f := range s.applyFuncs {
		db = f(db)
	}
	return db
}

func (s *BaseSearchEntity[T]) WithApply(a ApplyFunc) *BaseSearchEntity[T] {
	if a == nil {
		return s
	}
	s.applyFuncs = append(s.applyFuncs, a)
	return s
}

func (s *BaseSearchEntity[T]) WithEntity(e T) *BaseSearchEntity[T] {
	s.WithApply(func(d *gorm.DB) *gorm.DB {
		return d.Where(e)
	})
	return s
}

func (s *BaseSearchEntity[T]) WithID(id uuid.UUID) *BaseSearchEntity[T] {
	if id != uuid.Nil {
		s.WithApply(func(d *gorm.DB) *gorm.DB {
			return d.Where(&BaseEntity{
				ID: id,
			})
		})
	}
	return s
}

func (s *BaseSearchEntity[T]) WithLimit(limit interface{}) *BaseSearchEntity[T] {
	i := utils.InterfacetoInt(limit)
	s.WithApply(func(d *gorm.DB) *gorm.DB {
		return d.Limit(i)
	})
	return s
}

func (s *BaseSearchEntity[T]) WithDebug() *BaseSearchEntity[T] {
	s.WithApply(func(d *gorm.DB) *gorm.DB {
		return d.Debug()
	})
	return s
}

func (s *BaseSearchEntity[T]) WithOrder(column string, isDESC bool) *BaseSearchEntity[T] {
	s.WithApply(func(d *gorm.DB) *gorm.DB {
		return d.Order(clause.OrderByColumn{Column: clause.Column{Name: column}, Desc: isDESC})
	})
	return s
}

func (s *BaseSearchEntity[T]) WithPagination(limit, offset interface{}) *BaseSearchEntity[T] {
	return s.WithLimit(limit).WithOffset(offset)
}

func (s *BaseSearchEntity[T]) WithOffset(offset interface{}) *BaseSearchEntity[T] {
	i := utils.InterfacetoInt(offset)
	s.WithApply(func(d *gorm.DB) *gorm.DB {
		return d.Offset(i)
	})
	return s
}

func (s *BaseSearchEntity[T]) WithFields(fields ...string) *BaseSearchEntity[T] {
	return s.WithApply(func(d *gorm.DB) *gorm.DB {
		return d.Select(fields)
	})
}

type BaseRepository[T EntityPtr] struct {
	db *gorm.DB
}

func (s *BaseRepository[T]) SelectOne(a Applier) (T, error) {
	var v T
	if err := a.Apply(s.db).First(&v).Error; err != nil {
		return v, err
	}
	return v, nil
}

func (s *BaseRepository[T]) List(a Applier) ([]T, error) {
	var v []T
	if err := a.Apply(s.db).Find(&v).Error; err != nil {
		return v, err
	}
	return v, nil
}

func NewRepository[T EntityPtr]() RepositoryInterface[T] {
	if db == nil {
		panic("db not connected")
	}
	return &BaseRepository[T]{
		db: db,
	}
}
func NewSearch[T EntityPtr](entity T) *BaseSearchEntity[T] {
	return &BaseSearchEntity[T]{
		Entity: entity,
	}
}

func (s *BaseRepository[T]) CreateOne(value T) (T, error) {
	if err := s.db.Create(value).Error; err != nil {
		return value, err
	}
	return value, nil
}
