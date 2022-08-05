package generate

import (
	"fmt"
	"github.com/gookit/goutil/strutil"
	"github.com/hoisie/mustache"
	"github.com/ixfan/gofan/pkg/goc/template"
	"strings"
)

func Api(module, tableName string) error {
	if module == "" || tableName == "" {
		return fmt.Errorf("模块和表名不能为空")
	}
	goModule := GoModule()
	className := strutil.UpperFirst(strutil.CamelCase(tableName))
	table, err := GetTable(tableName)
	if err != nil {
		return err
	}
	//控制器
	controller := mustache.Render(template.Controller, map[string]interface{}{
		"module":       module,
		"goModule":     goModule,
		"className":    className,
		"lowerName":    strutil.CamelCase(tableName),
		"tableComment": table.Comment,
	})
	WriteFile("./internal/api/controller/"+tableName+".go", strings.TrimSpace(controller))
	//路由
	route := mustache.Render(template.Route, map[string]interface{}{
		"goModule":  goModule,
		"className": className,
		"lowerName": strutil.CamelCase(tableName),
	})
	WriteFile("./internal/api/route/"+tableName+".go", strings.TrimSpace(route))
	return nil
}
