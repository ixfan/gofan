package tools

type Pagination struct {
	Offset int
	Limit  int
}

func NewPagination(options map[string]interface{}) *Pagination {
	page := &Pagination{}
	pageSize := 10
	pageNumber := 1
	if ps, ok := options["PageSize"]; ok && ps.(int) > 0 {
		pageSize = ps.(int)
	}
	if pn, ok := options["Page"]; ok && pn.(int) > 0 {
		pageNumber = pn.(int)
	}
	page.Offset = pageSize * (pageNumber - 1)
	page.Limit = pageSize
	return page
}
