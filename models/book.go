package models

import "time"

type Book struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Author    string    `gorm:"not null" json:"author"`
	CreatedAt time.Time `gorm:"type:timestamp;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;autoCreateTime" json:"updated_at"`
}
