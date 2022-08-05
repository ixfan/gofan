package generate

import (
	"fmt"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/strutil"
)

func Domain(module, tableName string) error {
	if module == "" || tableName == "" {
		return fmt.Errorf("模块和表名不能为空")
	}
	columns, err := GetColumns(tableName)
	if err != nil {
		return err
	}
	table, err := GetTable(tableName)
	if err != nil {
		return err
	}
	//生成domain entity结构体
	ignores := []string{"created_by", "updated_by", "created_at", "updated_at", "deleted_at"}
	code := NewGenStruct("entity").SetTable(table).SetIgnores(ignores).GenerateColumn(columns) + (&GenStruct{}).SetTable(table).SetIgnores(ignores).GenerateUpdate(columns)
	WriteFile("./internal/domain/entity/"+tableName+".go", code)
	//生成model结构体
	code = NewGenStruct("model").SetTable(table).GenerateColumn(columns)
	WriteFile("./internal/infrastructure/model/"+tableName+".go", code)
	//生成应用层结构体
	//新增
	code = NewGenStruct("command").
		SetName("Create" + strutil.UpperFirst(strutil.CamelCase(table.Name))).
		SetComment("新增" + table.Comment).
		EnableLabelTag().
		SetIgnores([]string{"created_by", "updated_by", "created_at", "updated_at", "deleted_at", GetPrimaryKey(columns)}).
		GenerateColumn(columns)
	WriteFile("./internal/application/"+module+"/command/create_"+tableName+".go", code)
	//修改
	code = NewGenStruct("command").
		SetName("Update" + strutil.UpperFirst(strutil.CamelCase(table.Name))).
		SetComment("修改" + table.Comment).
		EnableLabelTag().
		SetIgnores([]string{"created_by", "updated_by", "created_at", "updated_at", "deleted_at"}).
		GenerateColumn(columns)
	WriteFile("./internal/application/"+module+"/command/update_"+tableName+".go", code)
	//删除
	code = NewGenStruct("command").
		SetName("Delete" + strutil.UpperFirst(strutil.CamelCase(table.Name))).
		SetComment("删除" + table.Comment).
		EnableLabelTag().
		GenerateQueryOne(columns)
	WriteFile("./internal/application/"+module+"/command/delete_"+tableName+".go", code)
	//查询
	code = NewGenStruct("query").
		SetName("Get" + strutil.UpperFirst(strutil.CamelCase(table.Name))).
		SetComment("获取" + table.Comment).
		EnableLabelTag().
		GenerateQueryOne(columns)
	WriteFile("./internal/application/"+module+"/query/get_"+tableName+".go", code)
	//列表
	code = NewGenStruct("query").
		SetName("List" + strutil.UpperFirst(strutil.CamelCase(table.Name))).
		SetComment(table.Comment + "列表").
		SetIgnores([]string{"created_by", "updated_by", "created_at", "updated_at", "deleted_at"}).
		GenerateList(columns)
	WriteFile("./internal/application/"+module+"/query/list_"+tableName+".go", code)
	return nil
}

func WriteFile(fileName string, content string) {
	if !fsutil.FileExists(fileName) {
		_, _ = fsutil.PutContents(fileName, content)
	}
}
