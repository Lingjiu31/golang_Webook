package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

// Insert 记录时间, 并且使用 grom 录入数据库
func (dao *UserDAO) Insert(cxt context.Context, user User) error {
	now := time.Now().UnixMilli()
	user.Utime = now
	user.Ctime = now
	return dao.db.WithContext(cxt).Create(&user).Error
}

// User 对应数据库表结构
type User struct {
	Id int64 `gorm:"primary_key,autoIncrement"`
	// 全部用户唯一
	Email    string `gorm:"unique"`
	Password string

	//创建时间,毫秒
	Ctime int64
	//更新时间,毫秒
	Utime int64
}
