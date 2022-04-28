package cockroach

import (
	"strings"

	"github.com/aqaurius6666/go-utils/go1.18/database"
	"gorm.io/gorm"
)

func ApplyPagination(defaultSearchModel database.DefaultSearchModel, db *gorm.DB) *gorm.DB {
	if defaultSearchModel.Limit > 0 {
		db = db.Limit(defaultSearchModel.Limit)
	}
	if defaultSearchModel.Skip > 0 {
		db = db.Offset(defaultSearchModel.Skip)
	}
	return db
}

func ApplySort(defaultSearchModel database.DefaultSearchModel, db *gorm.DB) *gorm.DB {
	if defaultSearchModel.OrderBy != "" {
		orderByList := strings.Fields(defaultSearchModel.OrderBy)
		orderTypeList := strings.Fields(defaultSearchModel.OrderType)
		for i, orderBy := range orderByList {
			orderType := "asc"
			if orderTypeList[i] != "asc" {
				orderType = "desc"
			}
			db = db.Order(orderBy + " " + orderType)
		}
	}
	return db
}
