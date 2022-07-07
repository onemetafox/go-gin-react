package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	id				int
	title			string
	content			string
	created_at		time.Time
	status			int
	UserID			int
	User			User `gorm:"constraint:OnUpate:CASCADE, OnDelete:SET NULL;"`
}