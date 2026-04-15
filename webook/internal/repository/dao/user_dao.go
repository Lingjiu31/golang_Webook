package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// 唯一索引冲突错误 1062
const uniqueConflictsErrNo uint16 = 1062

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
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
func (dao *UserDAO) Insert(ctx context.Context, user User) error {
	now := time.Now().UnixMilli()
	user.Utime = now
	user.Ctime = now
	// 数据库操作
	err := dao.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		var mysqlErr *mysql.MySQLError
		// 检验是否为数据库错误 并且查看是否为邮箱冲突错误
		if errors.As(err, &mysqlErr) && mysqlErr.Number == uniqueConflictsErrNo {
			// 邮箱冲突
			return ErrUserDuplicateEmail
		}
	}
	return err
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
