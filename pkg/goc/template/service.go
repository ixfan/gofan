package template

var Service = `
package service

import (
	"github.com/gookit/goutil/structs"
	"github.com/ixfan/gofan/pkg/database/orm"
	"github.com/ixfan/gofan/pkg/web"
	"github.com/jinzhu/copier"
	"{{goModule}}/internal/application/{{module}}/command"
	"{{goModule}}/internal/application/{{module}}/query"
	"{{goModule}}/internal/domain/entity"
	"{{goModule}}/internal/facade"
)

type {{className}}Service struct {
	context *web.Context
}

func New{{className}}Service(context *web.Context) *{{className}}Service {
	return &{{className}}Service{context: context}
}

//Create 新增
func (service *{{className}}Service) Create(create{{className}} *command.Create{{className}}) (interface{}, error) {
	if err := service.context.Validate(create{{className}}); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	{{lowerName}} := &entity.{{className}}{}
	_ = copier.Copy({{lowerName}}, create{{className}})
	{{lowerName}}, err := facade.{{className}}Repository(service.context).Save({{lowerName}})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return {{lowerName}}, nil
}

//Update 修改
func (service *{{className}}Service) Update(update{{className}} *command.Update{{className}}) (interface{}, error) {
	if err := service.context.Validate(update{{className}}); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	{{lowerName}}Repository := facade.{{className}}Repository(service.context)
	{{lowerName}}, err := {{lowerName}}Repository.FindById(update{{className}}.Id)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	{{lowerName}}.Update(structs.ToMap(update{{className}}))
	{{lowerName}}, err = facade.{{className}}Repository(service.context).Save({{lowerName}})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return {{lowerName}}, nil
}

//List 列表
func (service *{{className}}Service) List(list{{className}} *query.List{{className}}) (interface{}, error) {
	total, list, err := facade.{{className}}Repository(service.context).Find(&orm.Conditions{
		Pagination: &orm.Pagination{
			Page:     list{{className}}.Page,
			PageSize: list{{className}}.PageSize,
		},
	})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return web.NewPagination(total, list), nil
}

//Get 获取信息
func (service *{{className}}Service) Get(get{{className}} *query.Get{{className}}) (interface{}, error) {
	if err := service.context.Validate(get{{className}}); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	{{lowerName}}, err := facade.{{className}}Repository(service.context).FindOne(&orm.Conditions{
		Equal: map[string]interface{}{"id": get{{className}}.Id},
	})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return {{lowerName}}, nil
}

//Delete 删除
func (service *{{className}}Service) Delete(delete{{className}} *command.Delete{{className}}) (interface{}, error) {
	if err := service.context.Validate(delete{{className}}); err != nil {
		return nil, web.ThrowError(web.ArgError, err.Error())
	}
	{{lowerName}}, err := facade.{{className}}Repository(service.context).FindById(delete{{className}}.Id)
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	{{lowerName}}, err = facade.{{className}}Repository(service.context).Delete({{lowerName}})
	if err != nil {
		return nil, web.ThrowError(web.InternalServerError, err.Error())
	}
	return {{lowerName}}, nil
}

`
