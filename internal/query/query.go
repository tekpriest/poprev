package query

import "gorm.io/gorm"

func PaginateRows(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 20
		case page <= 0:
			page = 0
		}
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

func FilterProjectByStatus(status ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(status) > 0 && status[0] != "" {
			return db.Where("status = ?", status[0])
		}
		return db.Where("status = ?", "approved")
	}
}
