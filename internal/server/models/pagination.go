package models

type Pagination struct {
	Limit  int   `json:"limit"`
	Page   int   `json:"page"`
	Total  int64 `json:"total"`
	Offset int   `json:"offset"`
}

// NewPagination creates a new pagination object
func NewPagination(page, limit int) *Pagination {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit
	return &Pagination{
		Limit:  limit,
		Page:   page,
		Offset: offset,
	}
}

// SetTotal sets the total number of records
func (p *Pagination) SetTotal(total int64) {
	p.Total = total
}
