package template

var Repository = `
package repository

import (
	"github.com/ixfan/gofan/pkg/database/orm"
	"github.com/ixfan/gofan/pkg/tools"
	"github.com/ixfan/gofan/pkg/web"
	"github.com/jinzhu/copier"
	"{{goModule}}/internal/domain/entity"
	"{{goModule}}/internal/infrastructure/model"
)

type {{className}}Repository struct {
	context *web.Context
}

func New{{className}}Repository(context *web.Context) *{{className}}Repository {
	return &{{className}}Repository{context: context}
}

//Count 统计数量
func (repository *{{className}}Repository) Count(conditions *orm.Conditions) (int64, error) {
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.{{className}}{}, conditions)
	err := query.Count(&total).Error
	return total, err
}

//Save 保存
func (repository *{{className}}Repository) Save({{lowerName}} *entity.{{className}}) (*entity.{{className}}, error) {
	{{lowerName}}Model := &model.{{className}}{}
	err := copier.Copy({{lowerName}}Model, {{lowerName}})
	if err != nil {
		return {{lowerName}}, err
	}
	if {{lowerName}}Model.Id <= 0 {
		{{lowerName}}Model.Id, err = tools.NewSnowflakeId()
		if err != nil {
			return {{lowerName}}, err
		}
		{{lowerName}}.Id = {{lowerName}}Model.Id
		auth := repository.context.Auth()
		if auth != nil {
			{{lowerName}}Model.CreatedBy = auth.UserId
		}
		err = repository.context.Transaction().Context.Create({{lowerName}}Model).Error
	} else {
		auth := repository.context.Auth()
		if auth != nil {
			{{lowerName}}Model.UpdatedBy = auth.UserId
		}
		err = repository.context.Transaction().Context.Model({{lowerName}}Model).Updates({{lowerName}}Model).Error
	}
	return {{lowerName}}, err
}

//Find 获取列表
func (repository *{{className}}Repository) Find(conditions *orm.Conditions) (int64, []*entity.{{className}}, error) {
	list := make([]*entity.{{className}}, 0)
	var total int64
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.{{className}}{}, conditions)
	err := query.Count(&total).Error
	if err != nil {
		return total, list, err
	}
	err = query.Find(&list).Error
	return total, list, err
}

//FindOne 获取单条记录
func (repository *{{className}}Repository) FindOne(conditions *orm.Conditions) (*entity.{{className}}, error) {
	{{lowerName}} := &entity.{{className}}{}
	query := orm.AdvanceSearch(repository.context.Transaction(), &model.{{className}}{}, conditions)
	err := query.First({{lowerName}}).Error
	return {{lowerName}}, err
}

//FindById 根据ID查询
func (repository *{{className}}Repository) FindById(id int64) (*entity.{{className}}, error) {
	{{lowerName}}Model := &model.{{className}}{}
	{{lowerName}} := &entity.{{className}}{}
	err := orm.AdvanceSearch(repository.context.Transaction(), &model.{{className}}{}, &orm.Conditions{
		Equal: map[string]interface{}{"id": id},
	}).First({{lowerName}}Model).Error
	if err != nil {
		return {{lowerName}}, err
	}
	err = copier.Copy({{lowerName}}, {{lowerName}}Model)
	return {{lowerName}}, err
}

//Delete 删除
func (repository *{{className}}Repository) Delete({{lowerName}} *entity.{{className}}) (*entity.{{className}}, error) {
	{{lowerName}}Model := &model.{{className}}{}
	_ = copier.Copy({{lowerName}}Model, {{lowerName}})
	err := repository.context.Transaction().Context.Delete({{lowerName}}Model).Error
	return {{lowerName}}, err
}
`
