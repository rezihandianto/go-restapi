package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"type:varchar(255);not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;autoCreateTime" json:"updated_at"`
}
