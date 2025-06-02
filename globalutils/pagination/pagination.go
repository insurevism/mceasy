package pagination

type Page struct {
	Page  uint64 `json:"page" form:"page" query:"page"`
	Limit uint64 `json:"limit" form:"limit" query:"limit"`
}

type Offset struct {
	Offset uint64
	Limit  uint64
}

type Paginated[Data any] struct {
	Page      uint64 `json:"page"`
	PageTotal uint64 `json:"pageTotal"`
	Limit     uint64 `json:"limit"`
	Total     uint64 `json:"total"`
	Data      []Data `json:"data"`
}
