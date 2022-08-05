package generate

import (
	"fmt"
	"github.com/gookit/goutil/strutil"
	"github.com/hoisie/mustache"
	"github.com/ixfan/gofan/pkg/goc/template"
	"strings"
)

func Service(module, tableName string) error {
	if module == "" || tableName == "" {
		return fmt.Errorf("模块和表名不能为空")
	}
	goModule := GoModule()
	className := strutil.UpperFirst(strutil.CamelCase(tableName))
	table, err := GetTable(tableName)
	if err != nil {
		return err
	}
	service := mustache.Render(template.Service, map[string]interface{}{
		"module":       module,
		"goModule":     goModule,
		"className":    className,
		"lowerName":    strutil.CamelCase(tableName),
		"tableComment": table.Comment,
	})
	WriteFile("./internal/application/"+module+"/service/"+tableName+".go", strings.TrimSpace(service))
	return nil
}
