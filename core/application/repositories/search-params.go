package repositories

import "math"

type SortOpt string

const (
	ASC  SortOpt = "asc"
	DESC SortOpt = "desc"
)

type Sort struct {
	field   string
	sortOpt SortOpt
}

type SearchParams[Filter any] struct {
	page    int
	perPage int
	sorts   []Sort
	filters []Filter
}

func (searchParams *SearchParams[Filter]) SetPage(page *int) {
	searchParams.page = setInt(page, 1)
}

func (searchParams *SearchParams[Filter]) SetPerPage(perPage *int) {
	searchParams.page = setInt(perPage, 10)
}

func (searchParams *SearchParams[Filter]) SetSorts(sorts *[]Sort) {
	searchParams.sorts = setDefaultArray[Sort](sorts)
}

func (searchParams *SearchParams[Filter]) SetFilters(filters *[]Filter) {
	searchParams.filters = setDefaultArray[Filter](filters)
}

func (searchParams *SearchParams[Filter]) GetFilters() []Filter {
	return searchParams.filters
}

func (searchParams *SearchParams[Filter]) GetSorts() []Sort {
	return searchParams.sorts
}

func (searchParams *SearchParams[Filter]) GetPage() int {
	return searchParams.page
}

func (searchParams *SearchParams[Filter]) GetPerPage() int {
	return searchParams.perPage
}

func NewSearchParams[T any](page *int, perPage *int, sort *[]Sort, filters *[]T) SearchParams[T] {
	return SearchParams[T]{
		page:    setInt(page, 1),
		perPage: setInt(perPage, 10),
		sorts:   setDefaultArray(sort),
		filters: setDefaultArray(filters),
	}
}

func setInt(param *int, defaultValue int) int {
	if param == nil || math.IsNaN(float64(*param)) || *param <= 0 {
		return defaultValue
	}

	return *param
}

func setDefaultArray[T any](param *[]T) []T {
	if param != nil {
		return *param
	}

	return make([]T, 0)
}
