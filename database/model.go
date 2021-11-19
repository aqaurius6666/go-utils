package database

import (
	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id" json:"id"`
	UpdatedAt int64     `gorm:"autoUpdateTime:milli" json:"updatedAt"`
	CreatedAt int64     `gorm:"autoCreateTime:milli" json:"createdAt"`
	DeletedAt int64     `gorm:"autoDeleteTime:milli" json:"deletedAt"`
}

type DefaultSearchModel struct {
	Skip      int
	Limit     int
	OrderBy   string
	OrderType string
	Fields    []string
}

