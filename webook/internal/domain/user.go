package domain

import "time"

// User 领域对象
type User struct {
	Id        int64
	Email     string
	Password  string
	Name      string
	Birthday  string
	Biography string
	Ctime     time.Time
}
