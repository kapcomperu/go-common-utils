package base

type AbstractCriteria struct {
	Page int64  `schema:"page"`
	Limit  int64  `schema:"limit"`
	Sorting []SortType `schema:"sorting"`
}

type SortType struct {
	SortBy string `schema:"sortby"`
	SortDirection    bool   `schema:"sortDirection"`
}