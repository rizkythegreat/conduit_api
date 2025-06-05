package entity

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Slug      string         `json:"slug" gorm:"uniqueIndex;not null;size:100"`
	AuthorID  uint           `json:"author_id" gorm:"not null"`
	Author    User           `json:"author" gorm:"foreignKey:AuthorID"`
	Title     string         `json:"title" gorm:"not null;size:255"`
	Body      string         `json:"body" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
