package cockroach

import (
	"context"

	"github.com/aqaurius6666/go-utils/go1.18/database"
	"gorm.io/gorm"
)

func Select[T database.StructBaseEntityMustImpl](ctx context.Context, db *gorm.DB, search database.GormEntitySearcher[T]) (*T, error) {
	var v *T
	if err := search.ApplySearchById(db).WithContext(ctx).Select(v).Error; err != nil {
		return nil, err
	}
	return v, nil
}
