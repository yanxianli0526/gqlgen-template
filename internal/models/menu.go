package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	ID        uuid.UUID `gorm:"primaryKey;uniqueIndex;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt

	ItemName      string `gorm:"not null"`
	Price         int    `gorm:"not null"`
	IsStopSelling bool   `gorm:"not null"`
}
