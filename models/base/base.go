package base

import (
	"strconv"
	"time"
)

type MetaData struct {
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
	CreatedAt   TimestampTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt   TimestampTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	DeletedAt   TimestampTime `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
	CreatedBy   string    `json:"createdBy,omitempty" bson:"createdBy,omitempty"`
	UpdatedBy   string    `json:"updatedBy,omitempty" bson:"updatedBy,omitempty"`
	DeletedBy   string    `json:"deletedBy,omitempty" bson:"deletedBy,omitempty"`
	MadeAt      TimestampTime `json:"madeAt,omitempty" bson:"madeAt,omitempty"`
	/* END AUDIT */
}

func (m *MetaData) PopulateCommon(user string) {

	if m.CreatedRecord {
		m.CreatedAt.Time = time.Now()
		m.DeletedAt.Time = time.Time{}
		m.UpdatedAt.Time = time.Time{}
		m.CreatedBy = user
	}

	if m.UpdatedRecord {
		m.UpdatedAt.Time = time.Now()
		m.UpdatedBy = user
	}

	if m.DeletedRecord {
		m.DeletedAt.Time = time.Now()
		m.DeletedBy = user
	}

	m.MadeAt.Time = time.Now()

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

type TimestampTime struct {
	time.Time
}

///implement encoding.JSON.Marshaler interface
func (t *TimestampTime) MarshalJSON() ([]byte, error) {
	bin := make([]byte, 16)
	bin = strconv.AppendInt(bin[:0], t.Time.Unix(), 10)
	return bin, nil
}

func (t *TimestampTime) UnmarshalJSON(bin []byte) error {
	v, err := strconv.ParseInt(string(bin), 10, 64)
	if err != nil {
		return err
	}
	t.Time = time.Unix(v, 0)
	return nil
}