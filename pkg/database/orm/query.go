package orm

import (
	"fmt"
	"github.com/gookit/goutil/strutil"
	"gorm.io/gorm"
)

type Conditions struct {
	Like         map[string]interface{}   `json:"like"`
	Equal        map[string]interface{}   `json:"equal"`
	NotEqual     map[string]interface{}   `json:"notEqual"`
	Greater      map[string]interface{}   `json:"greater"`
	GreaterEqual map[string]interface{}   `json:"greaterEqual"`
	Less         map[string]interface{}   `json:"less"`
	LessEqual    map[string]interface{}   `json:"lessEqual"`
	In           map[string][]interface{} `json:"in"`
	NotIn        map[string][]interface{} `json:"notIn"`
	Range        map[string][]interface{} `json:"range"`
	Raw          string                   `json:"raw"`
	GroupBy      string                   `json:"groupBy"`
	OrderBy      string                   `json:"orderBy"`
	SelectRaw    string                   `json:"selectRaw"`
	Pagination   *Pagination              `json:"pagination"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

// DefaultPageSize 默认分页行数
var DefaultPageSize = 10

func AdvanceSearch(transaction *Transaction, model interface{}, condition *Conditions) *gorm.DB {
	tx := transaction.Context.Model(model)
	if condition.SelectRaw != "" {
		tx = tx.Select(condition.SelectRaw)
	}
	//模糊查询
	if len(condition.Like) > 0 {
		for key, item := range condition.Like {
			key = strutil.Snake(key)
			tx = tx.Where(key+" like ?", "%"+fmt.Sprintf("%v", item)+"%")
		}
	}
	//精确查询
	if len(condition.Equal) > 0 {
		for key, item := range condition.Equal {
			key = strutil.Snake(key)
			tx = tx.Where(key+" = ?", item)
		}
	}
	//不等于
	if len(condition.NotEqual) > 0 {
		for key, item := range condition.NotEqual {
			key = strutil.Snake(key)
			tx = tx.Where(key+" != ?", item)
		}
	}
	//大于
	if len(condition.Greater) > 0 {
		for key, item := range condition.Greater {
			key = strutil.Snake(key)
			tx = tx.Where(key+" > ?", item)
		}
	}
	//大于等于
	if len(condition.GreaterEqual) > 0 {
		for key, item := range condition.GreaterEqual {
			key = strutil.Snake(key)
			tx = tx.Where(key+" >= ?", item)
		}
	}
	//小于
	if len(condition.Less) > 0 {
		for key, item := range condition.Less {
			key = strutil.Snake(key)
			tx = tx.Where(key+" < ?", item)
		}
	}
	//小于等于
	if len(condition.LessEqual) > 0 {
		for key, item := range condition.LessEqual {
			key = strutil.Snake(key)
			tx = tx.Where(key+" <= ?", item)
		}
	}
	//范围
	if len(condition.Range) > 0 {
		for key, item := range condition.Range {
			key = strutil.Snake(key)
			if len(item) < 2 {
				continue
			}
			tx = tx.Where(key+" >= ? and "+key+" <= ?", item[0], item[1])
		}
	}
	if len(condition.In) > 0 {
		for key, item := range condition.In {
			key = strutil.Snake(key)
			if len(item) <= 0 {
				continue
			}
			tx = tx.Where(key+"in ?", item)
		}
	}
	if len(condition.NotIn) > 0 {
		for key, item := range condition.NotIn {
			key = strutil.Snake(key)
			if len(item) <= 0 {
				continue
			}
			tx = tx.Where(key+"not in ?", item)
		}
	}
	//SQL语句条件
	if condition.Raw != "" {
		tx = tx.Where(condition.Raw)
	}
	//分组
	if condition.GroupBy != "" {
		tx = tx.Group(condition.GroupBy)
	}
	//排序
	if condition.OrderBy != "" {
		tx = tx.Order(condition.OrderBy)
	}
	//分页
	if condition.Pagination != nil {
		page := condition.Pagination.Page
		if page <= 0 {
			page = 1
		}
		pageSize := condition.Pagination.PageSize
		if pageSize <= 0 {
			pageSize = DefaultPageSize
		}
		tx = tx.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	return tx
}
