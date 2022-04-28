package test

import (
	"testing"

	"github.com/aqaurius6666/go-utils/go1.18/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	DSN = "postgresql://anygonow:AnyGoNow!123@database-1.cpzkophtevbm.us-east-1.rds.amazonaws.com:5432/test"
)
var (
	_ database.BaseEntityMustImplStr  = UserEntity{}
	_ database.BaseEntitySetterGetter = (*UserEntity)(nil)
	_ database.TableNamer             = UserEntity{}
)

type UserEntity struct {
	database.BaseEntity
	Username *string `gorm:"column:username"`
}
type UserSearcher struct {
	database.BaseSearchEntity[*UserEntity]
}

var NewUserSearch = func() *database.BaseSearchEntity[*UserEntity] { return database.NewSearch(&UserEntity{}) }

func (s *UserEntity) GetBaseEntity() *database.BaseEntity {
	return &s.BaseEntity
}
func (s *UserEntity) SetBaseEntity(b *database.BaseEntity) {
	s.BaseEntity = *b
}
func (UserEntity) TableName() string {
	return "users"
}
func (s UserEntity) StrBaseEntityMustImpl() {}

func TestSelectOneById(t *testing.T) {
	var err error
	err = database.ConnectDB(DSN)
	assert.Nil(t, err, err)
	// err = database.Migrate(UserEntity{})
	// assert.Nil(t, err, err)
	repo := database.NewRepository[*UserEntity]()
	assert.NotNil(t, repo, "repo is nil")

	// username := "adad"
	// searcher := NewUserSearch().WithApply(func(d *gorm.DB) *gorm.DB {
	// 	return d.Where(&UserEntity{
	// 		Username: utils.StrPtr("usernam"),
	// 	})
	// })
	searcher := NewUserSearch().WithID(uuid.New())
	assert.NotNil(t, searcher, "seacher is nil")
	// user, err := repo.SelectOne(searcher.Apply)
	user, err := repo.SelectOne(searcher)
	assert.Nil(t, err, err)
	assert.NotNil(t, user, "user is nil")
}

func TestSelectOne(t *testing.T) {
	var err error
	err = database.ConnectDB(DSN)
	assert.Nil(t, err, err)
	// err = database.Migrate(UserEntity{})
	// assert.Nil(t, err, err)
	repo := database.NewRepository[*UserEntity]()
	assert.NotNil(t, repo, "repo is nil")

	// username := "adad"
	searcher := NewUserSearch().WithEntity(&UserEntity{}).WithLimit(1).WithOffset(3).WithFields()
	assert.NotNil(t, searcher, "seacher is nil")
	assert.NotNil(t, searcher.Entity, "searcher.Entity is nil")
	// searcher.WithEntity(&UserEntity{
	// 	BaseEntity: database.BaseEntity{
	// 		ID: uid,
	// 	},
	// })

	user, err := repo.SelectOne(searcher)
	assert.Nil(t, err, err)
	assert.NotNil(t, user, "user is nil")
}

func TestList(t *testing.T) {
	var err error
	err = database.ConnectDB(DSN)
	assert.Nil(t, err, err)
	// err = database.Migrate(UserEntity{})
	// assert.Nil(t, err, err)
	repo := database.NewRepository[*UserEntity]()
	assert.NotNil(t, repo, "repo is nil")

	// username := "adad"
	searcher := NewUserSearch().WithEntity(&UserEntity{})
	assert.NotNil(t, searcher, "seacher is nil")
	assert.NotNil(t, searcher.Entity, "searcher.Entity is nil")
	// searcher.WithEntity(&UserEntity{
	// 	BaseEntity: database.BaseEntity{
	// 		ID: uid,
	// 	},
	// })

	user, err := repo.List(searcher)
	assert.Nil(t, err, err)
	assert.NotNil(t, user, "user is nil")
}
