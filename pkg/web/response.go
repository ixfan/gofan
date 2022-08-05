package web

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Pagination struct {
	Total int64       `json:"total"`
	Rows  interface{} `json:"rows"`
}

func NewPagination(total int64, rows interface{}) *Pagination {
	return &Pagination{
		Total: total,
		Rows:  rows,
	}
}
