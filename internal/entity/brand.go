package entity

import "time"

type (
	// Brand entity.
	Brand struct {
		ID        uint16    `gorm:"primaryKey" json:"id"`
		NamaBrand string    `gorm:"not null;size:100;unique" json:"namaBrand"`
		CreatedAt time.Time `gorm:"not null" json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
)
