package dto

type ListRequest struct {
	QueryCriteria map[string][]string `json:"queryCriteria,omitempty"`
	SortBy        map[string]string   `json:"sortBy,omitempty"`
	PageSize      uint                `json:"pageSize,omitempty"`
	PageIndex     uint                `json:"pageIndex,omitempty"`
}

type ListResult struct {
	Items      interface{} `json:"items,omitempty"`
	TotalPages uint        `json:"totalPages,omitempty"`
	TotalItems uint64      `json:"totalItems,omitempty"`
}
