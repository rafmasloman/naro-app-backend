package entity

import "time"

type User struct {
	ID        string    `gorm:"column:id;primaryKey"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
