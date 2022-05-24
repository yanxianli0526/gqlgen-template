package models

import (
	"time"

	"github.com/google/uuid"
)

// User defines a user for the app
type User struct {
	ID        uuid.UUID `gorm:"primaryKey;uniqueIndex;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time
	DeletedAt *time.Time
	FirstName string
	LastName  string
	IdNumber  string
}
