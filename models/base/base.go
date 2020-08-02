package base

import "time"

type AbstractModel struct {
	/* START METADATA */
	TxId          string `json:"txId,omitempty"`
	RowNum        int    `json:"rowNum,omitempty"`
	NewEntry      bool   `json:"newEntry,omitempty,omitempty"`
	UpdatedRecord bool   `json:"updatedRecord,omitempty"`
	RemovedRecord bool   `json:"removedRecord,omitempty"`
	NewFile       bool   `json:"newFile,omitempty"`
	FileName      string `json:"fileName,omitempty"`
	ContentType   string `json:"contentType,omitempty"`
	/* END METADATA */

	StateId     int64     `json:"stateId,omitempty"`
	StateName   string    `json:"stateName,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	CreatedBy   string    `json:"createdBy,omitempty"`
	UpdatedBy   string    `json:"updatedBy,omitempty"`
	DeletedBy   string    `json:"deletedBy,omitempty"`
	FlagRemoved int       `json:"flagRemoved,omitempty"`
}

func (m *AbstractModel) PopulateCommon(user string) {

	if m.NewEntry {
		m.CreatedAt = time.Now()
		m.DeletedAt = time.Time{}
		m.UpdatedAt = time.Time{}
		m.StateId = 1
		m.CreatedBy = user
	}

	if m.UpdatedRecord {
		m.UpdatedAt = time.Now()
		m.UpdatedBy = user
	}

	if m.RemovedRecord {
		m.DeletedAt = time.Now()
		m.FlagRemoved = 1
		m.DeletedBy = user
	}

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
	Paging     Pageable     `json:"paging,omitempty"`
	StatusCode int           `json:"statusCode,omitempty"`
}
