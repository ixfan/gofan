package generate

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/gookit/goutil/strutil"
	"github.com/hoisie/mustache"
	"github.com/ixfan/gofan/pkg/goc/template"
	"strings"
)

func Repository(tableName string) error {
	if tableName == "" {
		return fmt.Errorf("表名不能为空")
	}
	goModule := GoModule()
	className := strutil.UpperFirst(strutil.CamelCase(tableName))
	//仓储接口
	importEntity := goModule + "/internal/domain/entity"
	code := jen.NewFile("interfaces")
	code.ImportName(goModule+"/internal/domain/entity", "entity")
	code.ImportName("github.com/ixfan/gofan/pkg/database/orm", "orm")
	code.Op("\n")
	code.Type().Id(className + "Repository").InterfaceFunc(func(group *jen.Group) {
		group.Id("Count").Call(jen.Op("*orm.Conditions")).Call(jen.Int64(), jen.Error()).Comment("统计记录数")
		group.Id("Save").Call(jen.Add(jen.Op("*").Qual(importEntity, className))).Call(jen.Op("*entity.").Id(className), jen.Error()).Comment("保存数据")
		group.Id("Find").Call(jen.Op("*orm.Conditions")).Call(jen.Int64(), jen.Index().Op("*entity.").Id(className), jen.Error()).Comment("列表")
		group.Id("FindOne").Call(jen.Op("*orm.Conditions")).Call(jen.Op("*entity.").Id(className), jen.Error()).Comment("查询单条记录")
		group.Id("FindById").Call(jen.Int64()).Call(jen.Op("*entity.").Id(className), jen.Error()).Comment("根据ID获取记录")
		group.Id("Delete").Call(jen.Add(jen.Op("*entity.").Id(className))).Call(jen.Op("*entity.").Id(className), jen.Error()).Comment("删除数据")
	})
	code.Op("\n")
	code.Type().Id(className + "Dao").Interface()
	WriteFile("./internal/domain/interfaces/"+tableName+".go", code.GoString())
	//仓储
	repository := mustache.Render(template.Repository, map[string]interface{}{
		"goModule":  goModule,
		"className": className,
		"lowerName": strutil.CamelCase(tableName),
	})
	WriteFile("./internal/infrastructure/repository/"+tableName+".go", strings.TrimSpace(repository))
	//dao
	dao := mustache.Render(template.Dao, map[string]interface{}{
		"className": className,
	})
	WriteFile("./internal/infrastructure/dao/"+tableName+".go", strings.TrimSpace(dao))
	//facade
	facade := mustache.Render(template.Facade, map[string]interface{}{
		"goModule":  goModule,
		"className": className,
	})
	WriteFile("./internal/facade/"+tableName+".go", strings.TrimSpace(facade))
	return nil
}
