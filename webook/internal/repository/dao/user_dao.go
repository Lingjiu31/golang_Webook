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
	ErrUserNotFound       = gorm.ErrRecordNotFound
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

// FindByEmail 根据邮箱寻找密码
func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return user, err
}

// UpDateUser 更新个人信息
func (dao *UserDAO) UpDateUser(ctx context.Context, user User) error {
	now := time.Now().UnixMilli()

	u := User{
		Name:      user.Name,
		Birthday:  user.Birthday,
		Biography: user.Biography,
		Utime:     now,
	}
	// 数据库操作
	return dao.db.WithContext(ctx).Model(&User{}).Where("id = ?", user.Id).Updates(&u).Error
}

// FindById 查找个人信息
func (dao *UserDAO) FindById(ctx context.Context, id int64) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return user, err
}

// User 对应数据库表结构
type User struct {
	Id       int64  `gorm:"primary_key,autoIncrement"` // 主键id
	Email    string `gorm:"unique"`                    // 全部用户唯一
	Password string

	Name      string
	Birthday  string
	Biography string
	// 创建时间, 毫秒
	Ctime int64
	// 更新时间, 毫秒
	Utime int64
}
