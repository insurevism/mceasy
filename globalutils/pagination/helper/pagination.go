package helper

import (
	"hokusai/globalutils/pagination"
	"math"
)

func OffsetFromPage(page *pagination.Page) pagination.Offset {

	if page == nil {
		return pagination.Offset{}
	}

	offset := (page.Page - 1) * page.Limit

	return pagination.Offset{Offset: offset, Limit: page.Limit}
}

func Paginator[Data any](page *pagination.Page, total uint64, data []Data) *pagination.Paginated[Data] {

	pagination := &pagination.Paginated[Data]{Data: data}

	pagination.Page = page.Page
	pagination.Limit = page.Limit
	pagination.Total = total

	if page.Limit != 0 {

		pagination.PageTotal = uint64(math.Ceil(float64(total) / float64(page.Limit)))
	}

	return pagination
}

// ReplacePaginatorData
// replace data and return new Copy
func ReplacePaginatorData[From any, To any](page *pagination.Paginated[From], data []To) *pagination.Paginated[To] {

	pagination := &pagination.Paginated[To]{Data: data}
	pagination.Page = page.Page
	pagination.Limit = page.Limit
	pagination.Total = page.Total
	pagination.PageTotal = page.PageTotal

	return pagination
}

func MergeDefaultPagination(page *pagination.Page, defaults *pagination.Page) (result *pagination.Page) {

	result = &pagination.Page{}
	result.Page = defaults.Page
	result.Limit = defaults.Limit

	if page == nil {
		page = &pagination.Page{}
	}

	if page.Limit > 0 {
		result.Limit = page.Limit
	}

	if page.Page > 0 {
		result.Page = page.Page
	}

	return
}
