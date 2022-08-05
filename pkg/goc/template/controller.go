package template

var Controller = `
package controller

import (
	"github.com/ixfan/gofan/pkg/web"
	"{{goModule}}/internal/application/{{module}}/command"
	"{{goModule}}/internal/application/{{module}}/query"
	"{{goModule}}/internal/application/{{module}}/service"
)

type {{className}}Controller struct {
	web.Controller
}

// Create 新增
// @Summary 创建{{tableComment}}
// @Tags {{tableComment}}信息
// @Accept json
// @Produce json
// @Router /api/v1/{{lowerName}} [post]
// @Param create{{className}} body command.Create{{className}} true "新建{{tableComment}}信息"
// @Success 200 {object} object{code=int,data=entity.{{className}}}
func (controller {{className}}Controller) Create(context *web.Context) {
	create{{className}} := &command.Create{{className}}{}
	_ = context.ShouldBindJSON(create{{className}})
	resp, err := service.New{{className}}Service(context).Create(create{{className}})
	controller.Response(context, resp, err)
}

// Update 修改
// @Summary 修改{{tableComment}}
// @Tags {{tableComment}}信息
// @Accept json
// @Produce json
// @Router /api/v1/{{lowerName}} [put]
// @Param update{{className}} body command.Update{{className}} true "修改{{tableComment}}信息"
// @Success 200 {object} object{code=int,data=entity.{{className}}}
func (controller {{className}}Controller) Update(context *web.Context) {
	update{{className}} := &command.Update{{className}}{}
	_ = context.ShouldBindJSON(update{{className}})
	resp, err := service.New{{className}}Service(context).Update(update{{className}})
	controller.Response(context, resp, err)
}

// List 列表
// @Summary {{tableComment}}列表
// @Tags {{tableComment}}信息
// @Accept json
// @Produce json
// @Router /api/v1/{{lowerName}}/list [post]
// @Param list{{className}} body query.List{{className}} true "修改{{tableComment}}信息"
// @Success 200 {object} object{code=int,data=object{total=int,list=[]entity.{{className}}}}
func (controller {{className}}Controller) List(context *web.Context) {
	list{{className}} := &query.List{{className}}{}
	_ = context.ShouldBindJSON(list{{className}})
	resp, err := service.New{{className}}Service(context).List(list{{className}})
	controller.Response(context, resp, err)
}

// Get 获取记录
// @Summary 获取{{tableComment}}
// @Tags {{tableComment}}信息
// @Accept json
// @Produce json
// @Router /api/v1/{{lowerName}} [get]
// @Param get{{className}} query query.Get{{className}} true "获取{{tableComment}}信息"
// @Success 200 {object} object{code=int,data=entity.{{className}}}
func (controller {{className}}Controller) Get(context *web.Context) {
	get{{className}} := &query.Get{{className}}{}
	_ = context.ShouldBind(get{{className}})
	resp, err := service.New{{className}}Service(context).Get(get{{className}})
	controller.Response(context, resp, err)
}

// Delete 删除
// @Summary 删除{{tableComment}}
// @Tags {{tableComment}}信息
// @Accept json
// @Produce json
// @Router /api/v1/{{lowerName}} [delete]
// @Param delete{{className}} body command.Delete{{className}} true "删除{{tableComment}}信息"
// @Success 200 {object} object{code=int,data=entity.{{className}}}
func (controller {{className}}Controller) Delete(context *web.Context) {
	delete{{className}} := &command.Delete{{className}}{}
	_ = context.ShouldBindJSON(delete{{className}})
	resp, err := service.New{{className}}Service(context).Delete(delete{{className}})
	controller.Response(context, resp, err)
}
`
