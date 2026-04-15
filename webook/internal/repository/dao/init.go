package dao

import "gorm.io/gorm"

// InitTable 在数据库中根据结构体建表(如果还没有这个表)
func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
