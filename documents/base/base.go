package base

import (
	"time"
)

type AbstractDocument struct {

	/* START METADATA */
	TxId          string `json:"txId,omitempty" bson:"txId,omitempty"`
	RowNum        int    `json:"rowNum,omitempty" bson:"rowNum,omitempty"`
	CreatedRecord bool   `json:"createdRecord,omitempty" bson:"createdRecord,omitempty"`
	UpdatedRecord bool   `json:"updatedRecord,omitempty" bson:"updatedRecord,omitempty"`
	DeletedRecord bool   `json:"deletedRecord,omitempty" bson:"deletedRecord,omitempty"`
	NewFile       bool   `json:"newFile,omitempty" bson:"newFile,omitempty"`
	FileName      string `json:"fileName,omitempty" bson:"fileName,omitempty"`
	ContentType   string `json:"contentType,omitempty" bson:"contentType,omitempty"`
	/* END METADATA */

	/* START AUDIT */
	StateId     string    `json:"stateId,omitempty" bson:"stateId,omitempty"`
	StateName   string    `json:"stateName,omitempty" bson:"stateName,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	DeletedAt   time.Time `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
	CreatedBy   string    `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	UpdatedBy   string    `json:"updatedBy,omitempty" bson:"updatedBy,omitempty"`
	DeletedBy   string    `json:"deletedBy,omitempty" bson:"deletedBy,omitempty"`
	MadeAt      time.Time `json:"madeAt,omitempty" bson:"madeAt,omitempty"`
	FlagRemoved int       `json:"flagRemoved,omitempty" bson:"flagRemoved,omitempty"`
	/* END AUDIT */
}

func (m *AbstractDocument) PopulateCommon(user string) {

	if m.CreatedRecord {
		m.CreatedAt = time.Now()
		m.DeletedAt = time.Time{}
		m.UpdatedAt = time.Time{}
		m.StateId = "01"
		m.CreatedBy = user
	}

	if m.UpdatedRecord {
		m.UpdatedAt = time.Now()
		m.UpdatedBy = user
	}

	if m.DeletedRecord {
		m.DeletedAt = time.Now()
		m.FlagRemoved = 1
		m.DeletedBy = user
	}

	m.MadeAt = time.Now()

}
