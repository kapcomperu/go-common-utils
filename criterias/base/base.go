package base

type AbstractCriteria struct {
	Page          int64  `schema:"page"`
	Limit         int64  `schema:"limit"`
	SortBy        string `schema:"sortBy"`
	SortDirection string `schema:"sortDirection"`
}