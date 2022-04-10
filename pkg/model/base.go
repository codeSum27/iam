package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID                uuid.UUID           `gorm:"size:36;primaryKey" json:"id`
	CreatedAt         time.Time           `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
    b.ID = uuid.NewV4()
	return
}

