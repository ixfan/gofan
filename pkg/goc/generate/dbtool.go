package generate

import (
	"fmt"
	"github.com/gookit/goutil/strutil"
	"github.com/ixfan/gofan/pkg/database/orm"
	"os"
	"strings"
)

type Column struct {
	Field   string      `json:"Field"`
	Type    string      `json:"Type"`
	Null    interface{} `json:"Null"`
	Key     string      `json:"Key"`
	Default interface{} `json:"Default"`
	Comment string      `json:"Comment"`
}

func (column Column) UpperCamelCaseField() string {
	return strutil.UpperFirst(strutil.CamelCase(column.Field))
}

func GetColumns(tableName string) ([]Column, error) {
	columns := make([]Column, 0)
	err := orm.Default().Context.Raw("show full columns from `" + tableName + "`").Find(&columns).Error
	return columns, err
}

type Table struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func GetTable(tableName string) (Table, error) {
	tableSchema := os.Getenv("mysql.database")
	tables := make([]Table, 0)
	err := orm.Default().Context.Raw("select TABLE_NAME as name,TABLE_COMMENT as comment from information_schema.Tables where TABLE_NAME=? and table_schema=?", tableName, tableSchema).Scan(&tables).Error
	if len(tables) > 0 && err == nil {
		return tables[0], nil
	}
	return Table{}, fmt.Errorf("未查询到该表")
}

//GetFiledType 获取字段类型
func GetFiledType(column Column) string {
	typeArr := strings.Split(column.Type, "(")
	switch typeArr[0] {
	case "int":
		return "int"
	case "integer":
		return "int"
	case "mediumint":
		return "int"
	case "bit":
		return "int"
	case "year":
		return "int"
	case "smallint":
		return "int"
	case "tinyint":
		return "int"
	case "bigint":
		return "int64"
	case "decimal":
		return "float32"
	case "double":
		return "float32"
	case "float":
		return "float32"
	case "real":
		return "float32"
	case "numeric":
		return "float32"
	case "timestamp":
		return "*time.Time"
	case "datetime":
		return "*time.Time"
	case "time":
		return "*time.Time"
	default:
		return "string"
	}
}

//GetPrimaryKey 获取主键
func GetPrimaryKey(columns []Column) string {
	for _, col := range columns {
		if col.Key == "PRI" {
			return col.Field
		}
	}
	return ""
}
