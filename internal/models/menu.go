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

	// Recipe        datatypes.JSON
	ItemName      string  `gorm:"not null"`
	Price         float64 `gorm:"not null"`
	IsStopSelling bool    `gorm:"not null"`
}

// func MarshalJsonType(a datatypes.JSON) graphql.Marshaler {
// 	return graphql.WriterFunc(func(w io.Writer) {
// 		data, _ := json.Marshal(a)
// 		io.WriteString(w, string(data))
// 	})
// }

// func UnmarshalJsonType(v interface{}) (datatypes.JSON, error) {
// 	a, ok := v.(datatypes.JSON)
// 	if !ok {
// 		return nil, errors.New("failed to cast to datatypes.JSON")
// 	}
// 	return a, nil
// }
