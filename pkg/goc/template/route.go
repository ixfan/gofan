package template

var Route = `
package route

import (
	"github.com/ixfan/gofan/pkg/web"
	"{{goModule}}/internal/api/controller"
)

func {{className}}() {
	{{lowerName}} := web.Default().Group("/api/v1/{{lowerName}}")
	{
		//新增
		{{lowerName}}.POST("/", func(context *gin.Context) {
			controller.{{className}}Controller{}.Create(web.NewContext(context))
		})
		//编辑
		{{lowerName}}.PUT("/", func(context *gin.Context) {
			controller.{{className}}Controller{}.Update(web.NewContext(context))
		})
		//查询
		{{lowerName}}.GET("/", func(context *gin.Context) {
			controller.{{className}}Controller{}.Get(web.NewContext(context))
		})
		//删除
		{{lowerName}}.DELETE("/", func(context *gin.Context) {
			controller.{{className}}Controller{}.Delete(web.NewContext(context))
		})
		//列表
		{{lowerName}}.POST("/list", func(context *gin.Context) {
			controller.{{className}}Controller{}.List(web.NewContext(context))
		})
	}
}
`
