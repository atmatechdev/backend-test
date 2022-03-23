package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Deleted     gorm.DeletedAt `gorm:"index" json:"deleted"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	Content     string         `json:"content" gorm:"not null"`
	CreatedById uint           `json:"created_by_id"`
	User        User           `gorm:"foreignKey:CreatedById" json:"-"`
}
