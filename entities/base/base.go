package base

import "time"

type AbstractEntity struct {
	/* START AUDIT */
	StateId     int64    `json:"stateId,omitempty" gorm:"size:255"`
	CreatedAt   time.Time `json:"createdAt,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	CreatedBy   string    `json:"createdBy,omitempty" gorm:"size:500"`
	UpdatedBy   string    `json:"updatedBy,omitempty" gorm:"size:500"`
	DeletedBy   string    `json:"deletedBy,omitempty" gorm:"size:500"`
	/* END AUDIT */
}
