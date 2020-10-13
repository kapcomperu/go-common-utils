package base

import (
	"time"
)

type MetaData struct {
	/* START METADATA */
	TxId          string `json:"txId,omitempty"`
	RowNum        int    `json:"rowNum,omitempty"`
	CreatedRecord bool   `json:"createdRecord,omitempty"`
	UpdatedRecord bool   `json:"updatedRecord,omitempty"`
	DeletedRecord bool   `json:"deletedRecord,omitempty"`
	NewFile       bool   `json:"newFile,omitempty"`
	FileName      string `json:"fileName,omitempty"`
	ContentType   string `json:"contentType,omitempty"`
	/* END METADATA */

	/* START AUDIT */
	StateId     string    `json:"stateId,omitempty"`
	StateName   string    `json:"stateName,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty" example:"2021-01-30T08:30:00Z"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty" example:"2021-01-30T08:30:00Z"`
	DeletedAt   time.Time `json:"deletedAt,omitempty" example:"2021-01-30T08:30:00Z"`
	CreatedBy   string    `json:"createdBy,omitempty"`
	UpdatedBy   string    `json:"updatedBy,omitempty"`
	DeletedBy   string    `json:"deletedBy,omitempty"`
	MadeAt      time.Time `json:"madeAt,omitempty" example:"2021-01-30T08:30:00Z"`
	/* END AUDIT */
}

func (m *MetaData) PopulateCommon(user string) {

	if m.CreatedRecord {
		m.CreatedAt = time.Now()
		m.DeletedAt = time.Time{}
		m.UpdatedAt = time.Time{}
		m.CreatedBy = user
	}

	if m.UpdatedRecord {
		m.UpdatedAt = time.Now()
		m.UpdatedBy = user
	}

	if m.DeletedRecord {
		m.DeletedAt = time.Now()
		m.DeletedBy = user
	}

	m.MadeAt = time.Now()

}
type AbstractModelCriteria struct {
	StateId       string
	SortColumn    string
	SortDirection string
	Page          int64
	Limit         int64
	NoPaging      bool
}

type Pageable struct {
	Total     int64 `json:"total,omitempty"`
	Page      int64 `json:"page,omitempty"`
	PerPage   int64 `json:"limit,omitempty"`
	Prev      int64 `json:"prev,omitempty"`
	Next      int64 `json:"next,omitempty"`
	TotalPage int64 `json:"totalPage,omitempty"`
}

type ResponseDataList struct {
	Data       []interface{} `json:"data"`
	Status     bool          `json:"status,omitempty"`
	StatusId   string        `json:"statusId,omitempty"`
	Messages   []string      `json:"messages,omitempty"`
	StatusCode int           `json:"statusCode,omitempty"`
}
type ResponseDataListPageable struct {
	Data       []interface{} `json:"data"`
	Status     bool          `json:"status,omitempty"`
	StatusId   string        `json:"statusId,omitempty"`
	Messages   []string      `json:"messages,omitempty"`
	Paging     *Pageable     `json:"paging,omitempty"`
	StatusCode int           `json:"statusCode,omitempty"`
}