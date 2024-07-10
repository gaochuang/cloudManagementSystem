package models

// Default values for page and page size
const (
	DefaultPage     = 1
	DefaultPageSize = 15
)

type PageResult struct {
	Items    interface{} `form:"items" json:"items"`       // Items holds the paginated data.
	Total    int64       `form:"total" json:"total"`       // Total is the total number of items.
	Page     int         `form:"page" json:"page"`         // Page is the current page number.
	PageSize int         `form:"pageSize" json:"pageSize"` // PageSize is the number of items per page.
}

func SetPageDefaults(p *PageResult) (page *PageResult, offset int) {
	if p.Page < 1 {
		p.Page = DefaultPage
	}
	if p.PageSize < 1 {
		p.PageSize = DefaultPageSize
	}

	// Calculate the offset for database query
	offset = p.PageSize * (p.Page - 1)

	return p, offset
}
